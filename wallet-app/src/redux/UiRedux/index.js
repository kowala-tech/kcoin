import { combineReducers } from "redux";
import modalReducer from "../../containers/ModalContainer/redux";
import messageReducer from "../../containers/MessageContainer/redux";

export default combineReducers({
	modal: modalReducer,
	message: messageReducer
});
