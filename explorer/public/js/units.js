var toKUSD = function(value) {
	return new BigNumber(value, 10).div('1000000000000000000').toString(10);
};
