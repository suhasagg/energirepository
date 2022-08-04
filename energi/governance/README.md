# Energi Governance
Here is the Energi Governance repository. This collection of smart contracts forms the basis of governance for the Energi DAO. These contracts are used by the [Energi Core Node](https://github.com/energicryptocurrency/energi) as a part of core consensus for the Energi blockchain. This repository may be used by any project seeking to interact with the Energi DAO.

## Building
Building this repository requires `nodejs-v12` and `yarn`. Other dependencies can be installed by yarn. The following steps will compile the governance contracts.
```sh
yarn install
yarn compile
```

## Testing
The tests can be run with yarn. The step `yarn run ganache` spawns a `ganache-cli` daemon which can be used for multiple runs of `yarn run test`.

```sh
yarn lint
yarn ganache
yarn test
```

## Deploying
These contracts are deployed via `truffle`. The **migrations/** directory specifies deployment of these contracts. You can deploy to `ganache` pretty easily:
```sh
yarn ganache
yarn truffle migrate -f 6
```

To deploy to the Energi testnet or mainnet, you will need a wallet to be able to pay the gas. You can set the environment variable `TRUFFLE_MNEMONIC` to the mnemonic phrase of a wallet that has some NRG to pay gas for deploying the contracts.
```sh
export TRUFFLE_MNEMONIC="dont share your seed phrase with anyone"
yarn truffle migrate --network testnet -f 6
yarn truffle migrate --network mainnet -f 6
```
