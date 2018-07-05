import {
	edgeWalletNamespace,
	getLocalEdgeAccount,
	getWallet,
	signAndSendTransaction,
} from "../../modules/edge";
import secureRandom from "secure-random";

export const createWallet = () => {
	return (dispatch) => {
		const privateKey = Buffer.from(secureRandom(32)).toString("hex");
		getLocalEdgeAccount().createWallet(
			edgeWalletNamespace,
			{ privateKey },
			(error, id) => {
				if (error) { return console.error(error); }
				dispatch(loadWallet(id));
			}
		);
	};
};

export function loadWallet(walletId) {
	return (dispatch) => {
		dispatch(setWalletLoading(true));
		const wallet = getWallet(walletId);
		const address = wallet.keys.address;
		const balance = wallet.getBalance();
		wallet.getTransactions({ startIndex: 1 }).then( (transactions) => {
			dispatch(replaceTransactions(transactions));
		});
		Promise.all([
			dispatch(setWalletId(walletId)),
			dispatch(setWalletAddress(address)),
			dispatch(setWalletBalance(balance))
		]).then( () => {
			dispatch(setWalletLoading(false));
		});
	};
}

export const sendTransaction = (walletId, toAddress, amountInWei) => {
	return () => {
		signAndSendTransaction(walletId, toAddress, amountInWei);
	};
};

export const setWalletLoading = (loading) => {
	return {
		type: "SET_WALLET_LOADING",
		loading
	};
};

export const setWalletId = (id) => {
	return {
		type: "SET_WALLET_ID",
		id
	};
};

export const setWalletBalance = (balance) => {
	return {
		type: "SET_WALLET_BALANCE",
		balance
	};
};

export const setWalletName = (name) => {
	return {
		type: "SET_WALLET_NAME",
		name
	};
};

export const addTransaction = (transaction) => {
	return {
		type: "ADD_TRANSACTION",
		transaction
	};
};

export const replaceTransactions = (transactions) => {
	return {
		type: "REPLACE_TRANSACTIONS",
		transactions
	};
};

export const setWalletAddress = (address) => {
	return {
		type: "SET_WALLET_ADDRESS",
		address
	};
};
