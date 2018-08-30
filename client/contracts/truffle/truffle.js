module.exports = {
    networks: {
        ganache: {
            host: '127.0.0.1',
            port: 8555,
            network_id: '*', // Match any network id
        },
        dev: {
            host: '127.0.0.1',
            port: 7545,
            network_id: '*',
        },
        kcoin: {
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
