/* eslint-disable import/default */

import React from "react";
import { render } from "react-dom";
import { AppContainer } from "react-hot-loader";
import configureStore, { history } from "./redux/configureStore";
import Root from "./components/Root";
import "./styles/styles.scss";
import injectTapEventPlugin from "react-tap-event-plugin";
require("./favicon.ico"); // Tell webpack to load favicon.ico
const store = configureStore();

render(
	<AppContainer>
		<Root store={store}
			history={history} />
	</AppContainer>,
	document.getElementById("app")
);

if (module.hot) {
	module.hot.accept("./components/Root", () => {
		const NewRoot = require("./components/Root").default;
		render(
			<AppContainer>
				<NewRoot store={store}
					history={history} />
			</AppContainer>,
			document.getElementById("app")
		);
	});
}

// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();
