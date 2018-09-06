'use strict';
var namehash = require('eth-ens-namehash');

var kns = function () {
    this.cachedDomainAddresses = [];
}

const ResolverAddr = "0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63";

kns.prototype.getKNSAddressFromDomain = function(domain, callback) {
    var _this = this;
    if (_this.cachedDomainAddresses[domain]) {
        return this.cachedDomainAddresses[domain];
    }

    let nameHashLib = require('eth-ens-namehash');

    let contractFunctionFullName = "addr(bytes32)";
    let funcSig = ethFuncs.getFunctionSignature(contractFunctionFullName);
    let input = '0x' + funcSig + ethUtil.solidityCoder.encodeParam("bytes32", nameHashLib.hash(domain))

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
                let domainAddr = ethUtil.solidityCoder.decodeParam("address", data.data.replace('0x', ''));
                _this.cachedDomainAddresses[domain] = domainAddr;

                callback(
                    {
                        error  : false,
                        message: "",
                        address: domainAddr,
                    },
                );
            }
        },
    );
};

module.exports = kns;