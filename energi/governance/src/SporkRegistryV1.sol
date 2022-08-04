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
import { IGovernedProxy } from "./IGovernedProxy.sol";
import { IUpgradeProposal } from "./IUpgradeProposal.sol";
import { ISporkRegistry } from "./ISporkRegistry.sol";
import { GenericProposalV1 } from "./GenericProposalV1.sol";

contract UpgradeProposalV1 is
    GenericProposalV1,
    IUpgradeProposal
{
    IGovernedContract public impl;
    address public creator;

    /**
     * C-tor
     *
     * @param _parent - actual parent, but not the caller
     * @param _impl - associated contract implementation
     * @param _mnregistry_proxy - IMasternodeRegistry proxy
     * @param _period - in seconds until deadline
     * @param _feePayer - the proposal initiator
     */
    constructor(
        address _parent,
        IGovernedContract _impl,
        IGovernedProxy _mnregistry_proxy,
        uint _period,
        address payable _feePayer
    )
        public
        GenericProposalV1(
            _mnregistry_proxy,
            QUORUM_MAJORITY,
            _period,
            _feePayer
        )
    {
        creator = msg.sender;
        parent = _parent;
        impl = _impl;
    }

    /**
     * Set fee amount by parent
     */
    function setFee() external payable {
        require(msg.sender == creator, "Only parent");
        // NOTE: make sure it correctly handles multiple calls
        fee_amount += msg.value;
    }
}

/**
 * Genesis hardcoded version of SporkRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract SporkRegistryV1 is
    GlobalConstants,
    ISporkRegistry,
    GovernedContract
{
    IGovernedProxy public mnregistry_proxy;

    // IGovernedContract
    //---------------------------------
    constructor(address _proxy, IGovernedProxy _mnregistry_proxy)
        public GovernedContract(_proxy)
    {
        mnregistry_proxy = _mnregistry_proxy;
    }

    // ISporkRegistry
    //---------------------------------
    function createUpgradeProposal(IGovernedContract _impl, uint _period, address payable _fee_payer)
        external payable
        returns (IUpgradeProposal proposal)
    {
        require(msg.value == FEE_UPGRADE_V1, "Invalid fee");
        require(_period >= PERIOD_UPGRADE_MIN, "Period min");
        require(_period <= PERIOD_UPGRADE_MAX, "Period max");

        proposal = new UpgradeProposalV1(
            msg.sender,
            _impl,
            mnregistry_proxy,
            _period,
            _fee_payer
        );

        proposal.setFee.value(msg.value)();
    }

    function consensusGasLimits()
        external view
        returns(uint callGas, uint xferGas)
    {
        callGas = 15000000;
        xferGas = 30000000;
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
