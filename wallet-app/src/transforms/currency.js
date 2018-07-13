import { BigNumber } from "bignumber.js";

const kusdAsWei = BigNumber(10).exponentiatedBy(18);

export const kusdToWei = (amountKusd) => {
	const normalizedKusd = BigNumber(String(amountKusd));
	return normalizedKusd.multipliedBy(kusdAsWei).toString();
};

export const weiToKusd = (amountWei) => {
	const normalizedWei = BigNumber(String(amountWei));
	return normalizedWei.dividedBy(kusdAsWei).toString();
};
