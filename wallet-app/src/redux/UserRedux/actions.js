import {
	createWallet,
	loadWallet,
	addTransaction,
	setWalletBalance
} from "../WalletRedux";
import { replace } from "react-router-redux";
import {
	getPrimaryWallet,
	setEdgeAccount,
	pinLogin,
	passwordLogin,
	createNewAccount
} from "../../modules/edge";

const loginCallbacks = (dispatch) => {
	return {
		callbacks: {
			onBalanceChanged: (walletId, currencyCode, balance) => {
				console.log( "Balance updated!");
				if (balance) { dispatch(setWalletBalance(balance.toString())); }
			},
			onTransactionsChanged: (walletId, txs) => {
				console.log( "Updated transaction!");
				dispatch(addTransaction(txs[0]));
			},
			onTxidsChanged: () => {
				console.log( "Updated transaction ids!");
			}
		}
	};
};

const loginLoading = () => {
	return {
		type: "LOGIN_LOADING"
	};
};

const loginSuccess = (account) => {
	return (dispatch) => {
		setEdgeAccount(account).then( () => {
			Promise.all([
				dispatch(setUsername(account.username)),
				dispatch(setAuthenticated(true)),
				dispatch(setPrimaryWallet(account)),
				dispatch(loginLoading()),
				dispatch(replace("/wallet"))
			]);
		});
	};
};

const loginError = (errorMessage) => {
	return {
		type: "LOGIN_ERROR",
		errorMessage
	};
};

const setPrimaryWallet = () => {
	return (dispatch) => {
		const primaryWallet = getPrimaryWallet();
		if (primaryWallet) {
			console.log("Primary wallet found with id " + primaryWallet.id);
			dispatch(loadWallet(primaryWallet.id));
			return;
		} else {
			console.error("No wallets found! Creating a new wallet.");
			dispatch(createWallet());
		}
	};
};

export const setUsername = (username) => {
	return {
		type: "SET_USERNAME",
		username
	};
};

export const setAuthenticated = (bool) => {
	return {
		type: "SET_AUTHENTICATED",
		bool
	};
};

export const loginWithPin = (username, pin, callback) => {
	return (dispatch) => {
		dispatch(loginLoading());
		pinLogin(
			username,
			pin,
			loginCallbacks(dispatch),
			(error) => { dispatch(loginError(error.message)); callback(); },
			(account) => { dispatch(loginSuccess(account)); }
		);
	};
};

export const loginWithPassword = (username, password) => {
	return (dispatch) => {
		dispatch(loginLoading());
		passwordLogin(
			username,
			password,
			loginCallbacks(dispatch),
			(error) => { dispatch(loginError(error.message)); },
			(account) => { dispatch(loginSuccess(account)); }
		);
	};
};

export const createAccount = (username, password, pin) => {
	return (dispatch) => {
		createNewAccount(
			username,
			password,
			pin,
			loginCallbacks(dispatch),
			(error) => { console.error({ message: "could not create wallet", error: error }); },
			(account) => { dispatch(loginSuccess(account)); }
		);
	};
};
