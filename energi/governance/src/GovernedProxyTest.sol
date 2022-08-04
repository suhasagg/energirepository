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

// solium-disable no-empty-blocks

import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IUpgradeProposal } from "./IUpgradeProposal.sol";
import { ISporkRegistry } from "./ISporkRegistry.sol";
import { IGovernedProxy, GovernedProxy } from "./GovernedProxy.sol";
import { StorageBase } from "./StorageBase.sol";
import { GovernedContractAutoProxy } from "./GovernedContractAutoProxy.sol";

contract MockContract is GovernedContract
{
    constructor(address _proxy) public GovernedContract(_proxy) {}
    function migrate(IGovernedContract) external {}
    function destroy(IGovernedContract new_impl) external {
        selfdestruct(address(new_impl));
    }
    function getAddress() external view returns (address) {
        return address(this);
    }
    function killStorage(StorageBase _storage) external {
        _storage.kill();
    }
    function testDrain() external {
        selfdestruct(msg.sender);
    }
    function testDrain(uint amount) external {
        msg.sender.transfer(amount);
    }
    function callProxy() external payable {
        address payable p = address(uint160(proxy));
        p.transfer(msg.value);
    }
    function () external payable {}
}

contract MockProxy is GovernedProxy
{
    constructor() public GovernedProxy(
        IGovernedContract(address(0)),
        new GovernedProxy(
            new MockSporkRegistry(address(0)),
            IGovernedProxy(address(0))
        )
    ) {}

    function setImpl(IGovernedContract _impl) external {
        impl = _impl;
    }

    function setSporkProxy(IGovernedProxy _proxy) external {
        spork_proxy = _proxy;
    }
}

contract MockSporkRegistry is MockContract, ISporkRegistry {
    constructor(address _proxy) public MockContract(_proxy) {}

    function createUpgradeProposal(IGovernedContract impl, uint, address payable)
        external payable
        returns (IUpgradeProposal)
    {
        return new MockProposal(msg.sender, impl);
    }

    function consensusGasLimits()
        external view
        returns(uint callGas, uint xferGas)
    {}
}

contract MockProposal is IUpgradeProposal {
    bool accepted;
    address public parent;
    IGovernedContract public impl;

    constructor(address _parent, IGovernedContract _impl) public {
        parent = _parent;
        impl = _impl;
    }

    function isAccepted() external view returns(bool) {
        return accepted;
    }
    function setAccepted() external {
        accepted = true;
    }
    function () external payable {}
    function created_block() external view returns(uint) {
        return 0;
    }
    function deadline() external view returns(uint) {
        return 0;
    }
    function fee_payer() external view returns(address payable) {
        return address(0);
    }
    function fee_amount() external view returns(uint) {
        return 0;
    }
    function accepted_weight() external view returns(uint) {
        return 0;
    }
    function rejected_weight() external view returns(uint) {
        return 0;
    }
    function total_weight() external view returns(uint) {
        return 0;
    }
    function quorum_weight() external view returns(uint) {
        return 0;
    }
    function isFinished() external view returns(bool) {
        return false;
    }
    function withdraw() external {}
    function destroy() external {
        require(msg.sender == parent, "Not Owner");
    }
    function collect() external {}
    function canVote(address) external view returns(bool) {
        return true;
    }
    function voteAccept() external {}
    function voteReject() external {}
    function setFee() external payable {}
}

// GovernedContractAutoProxy based implementation.

contract MockAutoContract is GovernedContractAutoProxy
{
    constructor() public GovernedContractAutoProxy(address(0)) {}
    function migrate(IGovernedContract) external {}
    function destroy(IGovernedContract new_impl) external {
        selfdestruct(address(new_impl));
    }
    function getAddress() external view returns (address) {
        return address(this);
    }
    function killStorage(StorageBase _storage) external {
        _storage.kill();
    }
    function testDrain() external {
        selfdestruct(msg.sender);
    }
    function testDrain(uint amount) external {
        msg.sender.transfer(amount);
    }
    function callProxy() external payable {
        address payable p = address(uint160(proxy));
        p.transfer(msg.value);
    }
    function () external payable {}
}

contract MockAutoProxy is GovernedProxy
{
    IGovernedContract impl;
    IGovernedProxy spork_proxy;

    constructor() public GovernedProxy(
        IGovernedContract(address(0)),
        new GovernedProxy(
            new MockAutoSporkRegistry(),
            IGovernedProxy(address(0))
        )
    ) {}

    function setImpl(IGovernedContract _impl) external {
        impl = _impl;
    }

    function setSporkProxy(IGovernedProxy _proxy) external {
        spork_proxy = _proxy;
    }
}

contract MockAutoSporkRegistry is MockAutoContract, ISporkRegistry {
    constructor() public MockAutoContract() {}

    function createUpgradeProposal(IGovernedContract impl, uint, address payable)
        external payable
        returns (IUpgradeProposal)
    {
        return new MockProposal(msg.sender, impl);
    }

    function consensusGasLimits()
        external view
        returns(uint callGas, uint xferGas)
    {}
}