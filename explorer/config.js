var fs = require('fs');

var config = {};

try {
    var network = process.env.NETWORK || 'testnet';
    var configFilename = __dirname + '/config.' + network + '.json';
    var configContents = fs.readFileSync(configFilename);
    config = JSON.parse(configContents);
    console.log(configFilename + ' found.');
}
catch (error) {
    if (error.code === 'ENOENT') {
        console.log('No config file found.');
    }
    else {
        throw error;
        process.exit(1);
    }
}


module.exports = config;
