module.exports = {
  networks: {
    ganache: {
      host: '127.0.0.1',
      port: 8555,
      network_id: '*', // Match any network id
    },
  },
  solc: {
    optimizer: {
      enabled: true,
      runs: 200,
    },
  },
};
