// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.16;
//pragma experimental SMTChecker;

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlacklistRegistry, IBlacklistProposal } from "./IBlacklistRegistry.sol";
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { ITreasury } from "./ITreasury.sol";
import { Gen2Migration } from "./Gen2Migration.sol";
import { GenericProposalV1 } from "./GenericProposalV1.sol";
import { StorageBase }  from "./StorageBase.sol";
import { NonReentrant } from "./NonReentrant.sol";

contract BlacklistProposalV1 is
    GenericProposalV1,
    IBlacklistProposal
{
    constructor(IGovernedProxy _mnregistry_proxy, address payable fee_payer)
        public
        GenericProposalV1(
            _mnregistry_proxy,
            10,
            1 weeks,
            fee_payer
        )
    // solium-disable-next-line no-empty-blocks
    {}

    function isObeyed()
        external view
        returns(bool)
    {
        if (isAccepted()) {
            return true;
        }

        uint accepted = accepted_weight;
        uint rejected = rejected_weight;

        if ((accepted > (rejected*2)) && (accepted > MN_COLLATERAL_MAX)) {
            return true;
        }

        return false;
    }
}

/**
 * A workaround for BlacklistRegistryV1 deploy-time gas consumption
 */
contract BlacklistV1ProposalCreator is
    StorageBase
{
    function create(IGovernedProxy mnregistry_proxy, address payable fee_payer)
        external payable
        requireOwner
        returns(IBlacklistProposal)
    {
        BlacklistProposalV1 proposal = new BlacklistProposalV1(
            mnregistry_proxy,
            fee_payer
        );

        proposal.setFee.value(msg.value)();
        return proposal;
    }
}


/**
 * Permanent storage of Blacklist Registry V1 data.
 */
contract StorageBlacklistRegistryV1 is
    StorageBase
{
    // NOTE: ABIEncoderV2 is not acceptable at the moment of development!

    struct Info {
        IBlacklistProposal enforce;
        IBlacklistProposal revoke;
        IBlacklistProposal drain;
        uint index;
    }

    mapping(address => Info) public address_info;
    address[] public address_list;

    function setEnforce(address addr, IBlacklistProposal proposal)
        external
        requireOwner
    {
        Info storage item = address_info[addr];
        assert(address(item.enforce) == address(0));

        item.enforce = proposal;
        item.index = address_list.length;
        address_list.push(addr);
    }

    function setRevoke(address addr, IBlacklistProposal proposal)
        external
        requireOwner
    {
        Info storage item = address_info[addr];

        assert(address(item.enforce) != address(0));

        item.revoke = proposal;
    }

    function setDrain(address addr, IBlacklistProposal proposal)
        external
        requireOwner
    {
        Info storage item = address_info[addr];

        assert(address(item.enforce) != address(0));

        item.drain = proposal;
    }

    function remove(address addr)
        external
        requireOwner
    {
        Info storage item = address_info[addr];
        assert(address(item.enforce) != address(0));

        // Ensure re-ordered index is updated
        address last = address_list[address_list.length - 1];
        address_info[last].index = item.index;

        // Move the last into the gap, NOOP on on match
        address_list[item.index] = last;
        address_list.pop();

        delete address_info[addr];
    }

    function addresses()
        external view
        returns(address[] memory result)
    {
        uint len = address_list.length;
        result = new address[](len);

        for (uint i = 0; i < len; ++i) {
            result[i] = address_list[i];
        }
    }
}


/**
 * Genesis hardcoded version of BlacklistRegistry.
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract BlacklistRegistryV1 is
    GovernedContract,
    NonReentrant,
    GlobalConstants,
    IBlacklistRegistry
{
    // Data for migration
    //---------------------------------
    BlacklistV1ProposalCreator public proposal_creator;
    StorageBlacklistRegistryV1 public v1storage;
    IGovernedProxy public mnregistry_proxy;
    Gen2Migration public migration;
    ITreasury public compensation_fund;
    address public EBI_signer;
    //---------------------------------

    constructor(
        address _proxy,
        IGovernedProxy _mnregistry_proxy,
        Gen2Migration _migration,
        ITreasury _compensation_fund,
        address _ebi_signer
    )
        public GovernedContract(_proxy)
    {
        proposal_creator = new BlacklistV1ProposalCreator();
        v1storage = new StorageBlacklistRegistryV1();
        mnregistry_proxy = _mnregistry_proxy;
        migration = _migration;
        compensation_fund = _compensation_fund;
        EBI_signer = _ebi_signer;
    }

    // IGovernedContract
    //---------------------------------
    function _destroy(IGovernedContract _newImpl) internal {
        v1storage.setOwner(_newImpl);
        proposal_creator.kill();
    }

    // IBlacklistRegistry
    //---------------------------------
    function proposals(address addr)
        external view
        returns(IBlacklistProposal enforce, IBlacklistProposal revoke, IBlacklistProposal drain)
    {
        (enforce, revoke, drain,) = v1storage.address_info(addr);
    }

    function _createProposal() internal returns(IBlacklistProposal) {
        // solium-disable-next-line security/no-low-level-calls
        (bool s, bytes memory r) = address(proposal_creator).delegatecall(
            abi.encodeWithSelector(
                proposal_creator.create.selector,
                mnregistry_proxy, _callerAddress())
        );
        require(s, string(r));
        return abi.decode(r, (IBlacklistProposal));
    }

    // solium-disable-next-line security/no-assign-params
    function _requireFee(uint fee) internal {
        if (_callerAddress() == EBI_signer) {
            fee = 0;
        }

        require(msg.value == fee, "Invalid fee");
    }

    function propose(address addr)
        external payable
        noReentry
        returns(IBlacklistProposal)
    {
        _requireFee(FEE_BLACKLIST_V1);

        StorageBlacklistRegistryV1 store = v1storage;
        (IBlacklistProposal enforce, IBlacklistProposal revoke, IBlacklistProposal drain,) = store.address_info(addr);

        // Cleanup old
        if (address(enforce) != address(0)) {
            if (address(revoke) != address(0)) {
                // assume enforced
                if (revoke.isAccepted()) {
                    enforce.destroy();
                    revoke.destroy();
                    if (address(drain) != address(0)) {
                        drain.destroy();
                    }
                    store.remove(addr);
                } else if (revoke.isFinished()) {
                    revert("Already active (1)");
                }
            } else if (enforce.isFinished() && !enforce.isAccepted()) {
                enforce.collect();
                // See below
                if (address(drain) != address(0)) {
                    drain.destroy();
                }
                store.remove(addr);
            } else {
                revert("Already active (2)");
            }
        }

        // Create new
        IBlacklistProposal proposal = _createProposal();

        store.setEnforce(addr, proposal);

        emit BlacklistProposal(addr, proposal);

        return proposal;
    }

    function proposeRevoke(address addr)
        external payable
        noReentry
        returns(IBlacklistProposal)
    {
        _requireFee(FEE_BLACKLIST_REVOKE_V1);

        StorageBlacklistRegistryV1 store = v1storage;
        (IBlacklistProposal enforce, IBlacklistProposal revoke,,) = store.address_info(addr);

        // Cleanup old
        require(address(enforce) != address(0), "No need (1)");

        if (address(revoke) != address(0)) {
            // assume enforced
            if (!revoke.isFinished()) {
                revert("Already active");
            } else if (!revoke.isAccepted()) {
                revoke.collect();
            }
        } else if (!enforce.isFinished()) {
            revert("Not applicable");
        } else if (!enforce.isAccepted()) {
            revert("No need (2)");
        }

        // Create new
        IBlacklistProposal proposal = _createProposal();

        store.setRevoke(addr, proposal);

        emit WhitelistProposal(addr, proposal);

        return proposal;
    }

    function proposeDrain(address addr)
        external payable
        noReentry
        returns(IBlacklistProposal)
    {
        _requireFee(FEE_BLACKLIST_DRAIN_V1);

        require(isBlacklisted(address(addr)), "Not blacklisted");

        StorageBlacklistRegistryV1 store = v1storage;
        (,, IBlacklistProposal drain,) = store.address_info(addr);

        if (address(drain) != address(0)) {
            if (drain.isAccepted()) {
                revert("Not need");
            } else if (drain.isFinished()) {
                drain.collect();
            } else {
                revert("Voting in progress");
            }
        }

        // Create new
        IBlacklistProposal proposal = _createProposal();

        store.setDrain(addr, proposal);

        emit DrainProposal(addr, proposal);

        return proposal;
    }

    function isBlacklisted(address addr)
        public view
        returns(bool)
    {
        StorageBlacklistRegistryV1 store = v1storage;
        (IBlacklistProposal enforce, IBlacklistProposal revoke,,) = store.address_info(addr);

        if ((address(revoke) != address(0)) && revoke.isAccepted()) {
            return false;
        }

        if ((address(enforce) != address(0)) && enforce.isObeyed()) {
            return true;
        }

        return false;
    }

    function isDrainable(address addr)
        public view
        returns(bool)
    {
        (IBlacklistProposal enforce, IBlacklistProposal revoke, IBlacklistProposal drain,) = v1storage.address_info(addr);

        if (address(enforce) == address(0)) {
            return false;
        }

        if (!enforce.isAccepted()) {
            return false;
        }

        if ((address(revoke) != address(0)) && revoke.isAccepted()) {
            return false;
        }

        if (address(drain) == address(0)) {
            return false;
        }

        return drain.isAccepted();
    }

    function collect(address addr)
        external
        noReentry
    {
        StorageBlacklistRegistryV1 store = v1storage;
        (IBlacklistProposal enforce, IBlacklistProposal revoke, IBlacklistProposal drain,) = store.address_info(addr);

        require(address(enforce) != address(0), "Nothing to collect");
        require(enforce.isFinished(), "Enforce voting in progress");

        if (!enforce.isAccepted()) {
            enforce.collect();
            store.remove(addr);
            return;
        }

        if (address(drain) != address(0)) {
            require(drain.isFinished(), "Drain voting in progress");

            if (drain.isAccepted()) {
                revert("Account must be drained");
            }

            drain.collect();
            store.setDrain(addr, IBlacklistProposal(address(0)));
            return;
        }

        if (address(revoke) != address(0)) {
            require(revoke.isFinished(), "Revoke voting in progress");

            if (revoke.isAccepted()) {
                enforce.destroy();
                revoke.destroy();
                assert(address(drain) == address(0));
                store.remove(addr);
            } else {
                revoke.collect();
                store.setRevoke(addr, IBlacklistProposal(address(0)));
            }
            return;
        }

        revert("No proposals ready to collect");
    }

    function drainMigration(uint item_id, bytes20 owner)
        external
        noReentry
    {
        require(isDrainable(address(owner)), "Not drainable");
        migration.blacklistClaim(item_id, owner);
        _onDrain(address(owner));
    }

    function enumerateAll()
        external view
        returns(address[] memory addresses)
    {
        return v1storage.addresses();
    }

    function enumerateBlocked()
        external view
        returns(address[] memory addresses)
    {
        addresses = v1storage.addresses();

        for (uint i = addresses.length; i-- > 0;) {
            if (!isBlacklisted(addresses[i])) {
                addresses[i] = address(0);
            }
        }
    }

    function enumerateDrainable()
        external view
        returns(address[] memory addresses)
    {
        addresses = v1storage.addresses();

        for (uint i = addresses.length; i-- > 0;) {
            if (!isDrainable(addresses[i])) {
                addresses[i] = address(0);
            }
        }
    }

    function onDrain(address addr)
        external
        noReentry
    {
        // solium-disable-next-line security/no-tx-origin
        require(tx.origin == proxy, "Not consensus");
        _onDrain(addr);
    }

    function _onDrain(address addr) internal {
        StorageBlacklistRegistryV1 store = v1storage;
        (IBlacklistProposal enforce, IBlacklistProposal revoke, IBlacklistProposal drain,) = store.address_info(addr);

        if (address(enforce) != address(0)) {
            enforce.destroy();

            if (address(revoke) != address(0)) {
                revoke.destroy();
            }

            if (address(drain) != address(0)) {
                drain.destroy();
            }

            store.remove(addr);
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
