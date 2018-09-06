'use strict';
var namehash = require('eth-ens-namehash');

var kns = function () {}

const ResolverAddr = "0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63";

kns.prototype.getKNSAddressFromDomain = function(domain, callback) {
    var nameHashLib = require('eth-ens-namehash');

    var contractFunctionFullName = "addr(bytes32)";
    var funcSig = ethFuncs.getFunctionSignature(contractFunctionFullName);
    var input = '0x' + funcSig + ethUtil.solidityCoder.encodeParam("bytes32", nameHashLib.hash(domain))

    ajaxReq.getEthCall(
        {
            to: ResolverAddr,
            data: input,
        },
        function (data) {
            if (data.error) {
                callback(
                    {
                        error  : true,
                        message: data.message,
                        address: null,
                    },
                );
            } else {
                callback(
                    {
                        error  : false,
                        message: "",
                        address: ethUtil.solidityCoder.decodeParam("address", data.data.replace('0x', '')),
                    },
                );
            }
        },
    );
};

module.exports = kns;