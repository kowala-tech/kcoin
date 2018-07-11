import Immutable from "seamless-immutable";

const INITIAL_STATE = Immutable({
	loading: false,
	error: false,
	errorMessage: "",
	username: "",
	authenticated: false
});

const loginReducer = (state = INITIAL_STATE, action) => {
	switch (action.type) {
	case "LOGIN_LOADING":
		return state.merge({
			loading: !state.loading
		});
	case "LOGIN_ERROR":
		return state.merge({
			loading: false,
			error: true,
			errorMessage: action.errorMessage
		});
	case "LOGIN_SUCCESS":
		return state.merge({
			loading: false,
			error: false,
			errorMessage: null,
			authenticated: true
		});
	case "SET_USERNAME":
		return state.merge({
			username: action.username
		});
	case "SET_AUTHENTICATED":
		return state.merge({
			authenticated: action.bool
		});
	default:
		return state;
	}
};

export default loginReducer;
