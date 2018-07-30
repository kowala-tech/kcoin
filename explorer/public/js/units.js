var toKUSD = function(value) {
    // 'ether': '1000000000000000000',
	return new BigNumber(value, 10).div('1000000000000000000').toString(10);
};
