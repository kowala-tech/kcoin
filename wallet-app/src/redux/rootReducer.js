import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";
import walletReducer from "../redux/WalletRedux";
import uiReducer from "../redux/UiRedux";
import userReducer from "../redux/UserRedux";
import { reducer as formReducer } from "redux-form";

const rootReducer = combineReducers({
	router: routerReducer,
	ui: uiReducer,
	wallet: walletReducer,
	user: userReducer,
	form: formReducer
});

export default rootReducer;
