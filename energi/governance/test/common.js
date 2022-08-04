// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

'use strict';

exports.mnregistry_config = [
    3,
    2,
    24*60*60,
    '10000000000000000000000', // 10,000 NRG
];
exports.mnregistry_config_v2 = [
    3, // RequireValidation
    3, // ValidationPeriods
    24*60*60, // CleanupPeriod
    '10000000000000000000000', // 10,000 NRG,
    1, // RewardsPerBlock: Default is 10 but changing this is resulting to lots of tests failures.
];
exports.superblock_cycles = 8;
exports.chain_id = 49797;
exports.migration_signer = '0x0000000000000000000000000000000012345678';
exports.cpp_signer = '0x2D0bc327d0843CAF6Fd9ae1eFaB0bF7196Fc2FC8';
exports.hf_finalization_period = 30;
exports.hf_signer = '0xF98eAeF47624cEE5C1328cdA0018381b5790A9e6';
exports.emergency_signer = '0x73E286b244c17F030F72e98C57FC83015a3C53Fd';
exports.default_address = '0x0000000000000000000000000000000000000000';
exports.zerofee_callopts = {
    gas: 250000, // real max is 500000
};
exports.mnreg_deploy_opts = {
    gas: 7000000,
};

exports.treasury_deploy_opts = {
    gas: 7000000,
};

exports.evt_last_block = {
    fromBlock: 'latest',
    toBlock: 'latest',
};

exports.stringifyBN = (web3, o) => {
    expect(o).is.not.undefined;
    for (let f in o) {
        const v = o[f];

        if (web3.utils.isBN(v)) {
            o[f] = v.toString();
        }
    }
    return o;
};

exports.moveTime = async (web3, seconds) => {
    expect(seconds).is.not.undefined;
    await new Promise((resolve, _) => {
        web3.currentProvider.send({
            jsonrpc: "2.0",
            method: "evm_increaseTime",
            params: [seconds],
            id: new Date().getSeconds()
        }, resolve);
    });
    await new Promise((resolve, _) => {
        web3.currentProvider.send({
            jsonrpc: "2.0",
            method: "evm_mine",
            params: [],
            id: new Date().getSeconds() + 1
        }, resolve);
    });
};

exports.govPreTests = (s) => {
    s.it('should refuse migrate() through s.proxy', async () => {
        try {
            await s.proxy_abi.migrate(s.fake.address, { from: s.accounts[0] });
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Good try/);
        }
    });

    s.it('should refuse destroy() through s.proxy', async () => {
        try {
            await s.proxy_abi.destroy(s.fake.address, { from: s.accounts[0] });
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Good try/);
        }
    });

    s.it('should refuse migrate() directly', async () => {
        try {
            await s.orig.migrate(s.fake.address, { from: s.accounts[0] });
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Not proxy/);
        }
    });

    s.it('should refuse destroy() directly', async () => {
        try {
            await s.orig.destroy(s.fake.address, { from: s.accounts[0] });
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Not proxy/);
        }
    });
};

exports.govPostTests = (s) => {
    const MockProposal = s.artifacts.require('MockProposal');

    s.it('should refuse to accept funds', async () => {
        try {
            await s.token_abi.send(s.web3.utils.toWei('1', "ether"));
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Not supported/);
        }
    });

    if ('storage' in s) s.it('should refuse to accept funds to storage', async () => {
        try {
            await s.storage.send(s.web3.utils.toWei('1', "ether"));
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /revert/);
        }
    });

    if ('storage' in s) s.it('should refuse kill() storage', async () => {
        try {
            await s.storage.kill();
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Not owner/);
        }
    });

    if ('storage' in s) s.it('should refuse setOwner() on storage', async () => {
        try {
            await s.storage.setOwner(s.proxy.address);
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /Not owner/);
        }
    });

    s.it('should destroy() after upgrade', async () => {
        const orig_balance = await s.web3.eth.getBalance(s.orig.address)
        const { logs } = await s.proxy.proposeUpgrade(
            s.fake.address, 0,
            { from: s.accounts[0], value: '1' });

        s.assert.equal(logs.length, 1);
        const proposal = await MockProposal.at(logs[0].args['1']);

        await proposal.setAccepted();
        await s.proxy.upgrade(proposal.address);

        const fake_balance = await s.web3.eth.getBalance(s.fake.address)
        s.assert.equal(orig_balance.valueOf(), fake_balance.valueOf());

        try {
            await s.orig.proxy();
            s.assert.fail("It must fail");
        } catch (e) {
            s.assert.match(e.message, /did it run Out of Gas/);
        }
    });

    if ('storage' in s) s.it('should transfer storage & allow to kill() it', async () => {
        await s.fake.killStorage(s.storage.address);
    });
};
