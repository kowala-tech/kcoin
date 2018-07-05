export const normalizeAddress = (address) => {
	return address && !address.startsWith("0x") ? "0x".concat(address) : address;
};

export const unprefixAddress = (address) => {
	return address.replace(/^0x/, "");
};

export const collapseAddress = (address) => {
	const unprefixedAddress = unprefixAddress(address);
	return unprefixedAddress.slice(0,4) + "..." + unprefixedAddress.slice(-5);
};
