const HDWalletProvider = require("@truffle/hdwallet-provider");
const mnemonicPhrase = process.env.TRUFFLE_MNEMONIC;
const mochaReporters = process.env.MOCHA_REPORTERS || "spec" // to get test output XML use MOCHA_REPORTERS="mocha-junit-reporter, spec"
const providerOrUrl = process.env.PROVIDER_OR_URL;

module.exports = {
    contracts_directory: './src',
    contracts_build_directory: './build',
    migrations_directory: './migrations',
    test_directory: './test',
    test_file_extension_regexp: /.*\.spec\.js$/,
    verboseRpc: false,
    mocha: {
        spec: './test/*.spec.js',
        useColors: true,
        reporter: "mocha-multi-reporters",
        reporterOptions: {
            "reporterEnabled": mochaReporters,
            "mochaJunitReporterReporterOptions": {
                "mochaFile": ".test-results.xml"
            }
        }
    },
    networks: {
        development: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "*" // Match any network id
        },
        testnet: {
            provider: () =>
                new HDWalletProvider({
                    mnemonic: {
                        phrase: mnemonicPhrase
                    },
                    providerOrUrl: providerOrUrl || "https://nodeapi.test.energi.network",
                    derivationPath: "m/44'/49797'/0'/0/"
                }),
            network_id: "49797",
            gas: 30000000
        },
        mainnet: {
            provider: () =>
                new HDWalletProvider({
                    mnemonic: {
                        phrase: mnemonicPhrase
                    },
                    providerOrUrl: providerOrUrl || "https://nodeapi.energi.network",
                    derivationPath: "m/44'/39797'/0'/0/"
                }),
            network_id: "39797",
            gas: 30000000
        }
    },
    compilers: {
        solc: {
            version: '0.5.16',
            evmVersion: 'petersburg',
            optimizer: {
                enabled: true,
                runs: 9999999999,
            }
        }
    }
}
