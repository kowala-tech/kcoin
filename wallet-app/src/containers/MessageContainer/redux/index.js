import Immutable from "seamless-immutable";

const INITIAL_STATE = Immutable({
	open: false,
	text: ""
});

const messageReducer = (state = INITIAL_STATE, action) => {
	switch (action.type) {
	case "OPEN_MESSAGE":
		return state.merge({
			open: true,
			text: action.text
		});
	case "CLOSE_MESSAGE":
		return state.merge({
			open: false
		});
	default:
		return state;
	}
};

export default messageReducer;
