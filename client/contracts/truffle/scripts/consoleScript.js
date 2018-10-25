var multisigcreatorAddr = '0xFF9DFBD395cD1C4a4F23C16aa8a5c44109Bc17DF';
var multisigAddr = '0x0e5d0Fd336650E663C710EF420F85Fb081E21415';
var PublicResolverABI = [{'constant':true,'inputs':[{'name':'interfaceID','type':'bytes4'}],'name':'supportsInterface','outputs':[{'name':'','type':'bool'}],'payable':false,'stateMutability':'pure','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'key','type':'string'},{'name':'value','type':'string'}],'name':'setText','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':true,'inputs':[],'name':'initialized','outputs':[{'name':'','type':'bool'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'},{'name':'contentTypes','type':'uint256'}],'name':'ABI','outputs':[{'name':'contentType','type':'uint256'},{'name':'data','type':'bytes'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'x','type':'bytes32'},{'name':'y','type':'bytes32'}],'name':'setPubkey','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'}],'name':'content','outputs':[{'name':'','type':'bytes32'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'}],'name':'addr','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'},{'name':'key','type':'string'}],'name':'text','outputs':[{'name':'','type':'string'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'contentType','type':'uint256'},{'name':'data','type':'bytes'}],'name':'setABI','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'}],'name':'name','outputs':[{'name':'','type':'string'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'name','type':'string'}],'name':'setName','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'hash','type':'bytes'}],'name':'setMultihash','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'hash','type':'bytes32'}],'name':'setContent','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'knsAddr','type':'address'}],'name':'initialize','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'}],'name':'pubkey','outputs':[{'name':'x','type':'bytes32'},{'name':'y','type':'bytes32'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'node','type':'bytes32'},{'name':'addr','type':'address'}],'name':'setAddr','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':true,'inputs':[{'name':'node','type':'bytes32'}],'name':'multihash','outputs':[{'name':'','type':'bytes'}],'payable':false,'stateMutability':'view','type':'function'},{'inputs':[{'name':'knsAddr','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'constructor'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'a','type':'address'}],'name':'AddrChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'hash','type':'bytes32'}],'name':'ContentChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'name','type':'string'}],'name':'NameChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':true,'name':'contentType','type':'uint256'}],'name':'ABIChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'x','type':'bytes32'},{'indexed':false,'name':'y','type':'bytes32'}],'name':'PubkeyChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'indexedKey','type':'string'},{'indexed':false,'name':'key','type':'string'}],'name':'TextChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':true,'name':'node','type':'bytes32'},{'indexed':false,'name':'hash','type':'bytes'}],'name':'MultihashChanged','type':'event'}];
var AdminUpgradabilityProxyAbi = [{'inputs':[{'name':'_implementation','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'constructor'},{'payable':true,'stateMutability':'payable','type':'fallback'},{'anonymous':false,'inputs':[{'indexed':false,'name':'previousAdmin','type':'address'},{'indexed':false,'name':'newAdmin','type':'address'}],'name':'AdminChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':false,'name':'implementation','type':'address'}],'name':'Upgraded','type':'event'},{'constant':true,'inputs':[],'name':'admin','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[],'name':'implementation','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'newAdmin','type':'address'}],'name':'changeAdmin','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'}],'name':'upgradeTo','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'},{'name':'data','type':'bytes'}],'name':'upgradeToAndCall','outputs':[],'payable':true,'stateMutability':'payable','type':'function'}];
// var publicResolverAddr = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

var knsRegistryAddr = '0x4195B06a6e4d5bEDDe15165e01A64F324f03D5d1';
var knsRegistrarAddr = '0xe1adb6075619f52Fc00BDD50dEf1B754b9e7bd17';
var knsResolverAddr = '0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63';

var domain1 = 'systemvars.kowala';
var domain2 = 'oraclemgr.kowala';
var domain3 = 'validatormgr.kowala';
var domain4 = 'miningtoken.kowala';
var domain5 = 'stability.kowala';

namehash = function namehash(name){var node = '0x0000000000000000000000000000000000000000000000000000000000000000';if (name !== '') {var labels = name.split(".");for(var i = labels.length - 1; i >= 0; i--) {node = web3.sha3(node + web3.sha3(labels[i]).slice(2), {encoding: 'hex'});}}return node.toString();}

upgradeAdmin = function() {
  try {
    var publicResolver = web3.eth.contract(PublicResolverABI).at(knsResolverAddr);
    var sysvarAddr =  publicResolver.addr(namehash(domain1));
    console.log(sysvarAddr);
    var oraclemgrAddr = publicResolver.addr(namehash(domain2));
    console.log(oraclemgrAddr);
    var validatormgrAdrr = publicResolver.addr(namehash(domain3));
    console.log(validatormgrAdrr);
    var miningtokenAddr = publicResolver.addr(namehash(domain4));
    console.log(miningtokenAddr);
    var stabilityAddr = publicResolver.addr(namehash(domain5));
    console.log(stabilityAddr);

    var knsRegistryProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(knsRegistryAddr);
    var knsResolverProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(knsResolverAddr);
    var knsRegistrarProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(knsRegistrarAddr);
    
    var sysvarProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(sysvarAddr);
    var oraclemgrProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(oraclemgrAddr);
    var validatormgrProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(validatormgrAdrr);
    var miningtokenProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(miningtokenAddr);
    var stabilityProxy = web3.eth.contract(AdminUpgradabilityProxyAbi).at(stabilityAddr);

    sysvarProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(sysvarProxy.admin({ from: multisigAddr }));
    oraclemgrProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(oraclemgrProxy.admin({ from: multisigAddr }));
    validatormgrProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(validatormgrProxy.admin({ from: multisigAddr }));
    miningtokenProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(miningtokenProxy.admin({ from: multisigAddr }));
    stabilityProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(stabilityProxy.admin({ from: multisigAddr }));
    knsRegistryProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(knsRegistryProxy.admin({ from: multisigAddr }));
    knsResolverProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
    console.log(knsResolverProxy.admin({ from: multisigAddr }));
    knsRegistrarProxy.changeAdmin(multisigAddr, { from: multisigcreatorAddr });
  } catch (err) { console.log(err); }
}
