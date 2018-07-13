import { makeContext } from "edge-core-js";
import { kowalaCurrencyPluginFactory } from "edge-currency-kowala";

const edgeWalletName = KOWALA_NETWORK;
export const edgeWalletNamespace = "wallet:" + edgeWalletName;

export const edge = makeContext({
	apiKey: "e6eee331afb0385b6a6223719802fcfd00fc2331",
	plugins: [ kowalaCurrencyPluginFactory ],
	appId: edgeWalletName
});

export const storeEdgeOnWindow = async () => {
	return window.abcui = edge;
};

export const getLocalUsername = async () => {
	const usernameList = await window.abcui.usernameList();
	return usernameList[0];
};

export const logout = async () => {
	return window.abcui.abcAccount.logout();
};

export const deleteLocalAccount = async (username) => {
	return window.abcui.deleteLocalAccount(username);
};

export const pinLogin = async (username, pin, loginCallbacks, errorCallback, successCallback) => {
	return await window.abcui.loginWithPIN(
		username,
		String(pin),
		loginCallbacks,
		function (error, account) {
			if (error) {
				errorCallback(error);
			} else {
				successCallback(account);
			}
		}
	);
};

export const passwordLogin = async (username, password, loginCallbacks, errorCallback, successCallback) => {
	return await window.abcui.loginWithPassword(
		username,
		password,
		loginCallbacks,
		function (error, account) {
			if (error) {
				errorCallback(error);
			} else {
				successCallback(account);
			}
		}
	);
};

export const usernameAvailable = async (username) => {
	const valid = username.length > 4;
	const available = await window.abcui.usernameAvailable(username);
	return valid && available;
};

export const createNewAccount = async (username, password, pin, loginCallbacks, errorCallback, successCallback) => {
	return await window.abcui.createAccount(
		username,
		password,
		String(pin),
		loginCallbacks,
		function (error, account) {
			if (error) {
				errorCallback(error);
			} else {
				successCallback(account);
			}
		}
	);
};

const timeout = ms => new Promise(res => setTimeout(res, ms));

export const setEdgeAccount = async (account) => {
	await timeout(500); // edge is slow...so we wait
	window.abcui.abcAccount = account;
	return account;
};

export const getLocalEdgeAccount = () => {
	try {
		return window.abcui.abcAccount;
	}
	catch(error) {
		console.error("No local account found!");
	}
};

export const getWallet = (id) => {
	console.log("Getting wallet " + id);
	const wallets = getLocalEdgeAccount().currencyWallets;
	return wallets[id];
};

export const getPrimaryWallet = () => {
	const account = getLocalEdgeAccount();
	console.log("Getting wallets for account " + account.id);
	const primaryWallet = account.getFirstWalletInfo(edgeWalletNamespace);
	return primaryWallet;
};

export const signAndSendTransaction = async (walletId, toAddress, amountInWei) => {
	try {
		const wallet = getWallet(walletId);
		const spendParams = {
			networkFeeOption: "low",
			currencyCode: "KUSD",
			spendTargets: [
				{
					publicAddress: toAddress,
					nativeAmount: amountInWei
				}
			]
		};
		let transaction = await wallet.makeSpend(spendParams);
		transaction = await wallet.signTx(transaction);
		transaction = await wallet.broadcastTx(transaction);
		transaction = await wallet.saveTx(transaction);
		console.log("Sent transaction with ID = " + transaction.txid);
	} catch (error) {
		console.log(error);
	}
};

export default edge;
