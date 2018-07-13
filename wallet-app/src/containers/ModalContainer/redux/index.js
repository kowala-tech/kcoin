import Immutable from "seamless-immutable";

const INITIAL_STATE = Immutable({
	open: false,
	type: null
});

const modalReducer = (state = INITIAL_STATE, action) => {
	switch (action.type) {
	case "OPEN_MODAL":
		return state.merge({
			open: true,
			type: action.modalType
		});
	case "CLOSE_MODAL":
		return state.merge({
			open: false
		});
	default:
		return state;
	}
};

export default modalReducer;
