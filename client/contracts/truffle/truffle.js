module.exports = {
    networks: {
        ganache: {
            host: '127.0.0.1',
            port: 8545,
            network_id: '*', // Match any network id
        },
        kcoin_test: {
            host: '0.0.0.0',
            port: 30503,
            network_id: '*',
            gas: 2000000,
        },
        kcoin_main: {
            host: '127.0.0.1',
            port: 30503,
            network_id: '*',
        },
        coverage: {
            host: '127.0.0.1',
            network_id: '*',
            port: 8545,
            gasPrice: 0x1,
          },
    },
    solc: {
        optimizer: {
            enabled: true,
            runs: 200,
        },
    },
};
