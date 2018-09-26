var keythereum = require("keythereum");
var datadir = "/Users/juliusbrain/Work/Kowala/testnet/";
var address= "2429f4aa5cf9d23fea0961780ffb4ff8916a26a0";
const password = "test";

var keyObject = keythereum.importFromFile(address, datadir);
var privateKey = keythereum.recover(password, keyObject);
console.log(privateKey.toString('hex'));
