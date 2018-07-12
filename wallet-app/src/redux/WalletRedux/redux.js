import Immutable from "seamless-immutable";

const INITIAL_STATE = Immutable({
	loading: false,
	address: null,
	balance: "0",
	name: null,
	id: null,
	transactions: []
});

const walletReducer = (state = INITIAL_STATE, action) => {
	switch (action.type) {
	case "SET_WALLET_LOADING":
		return state.merge({
			loading: action.loading
		});
	case "SET_WALLET_ID":
		return state.merge({
			id: action.id
		});
	case "SET_WALLET_ADDRESS":
		return state.merge({
			address: action.address
		});
	case "SET_WALLET_NAME":
		return state.merge({
			name: action.name
		});
	case "SET_WALLET_BALANCE":
		return state.merge({
			balance: action.balance
		});
	case "ADD_TRANSACTION":
		return state.merge({
			transactions: state.transactions.concat(action.transaction)
		});
	case "REPLACE_TRANSACTIONS":
		return state.merge({
			transactions: action.transactions
		});
	default:
		return state;
	}
};

export default walletReducer;
