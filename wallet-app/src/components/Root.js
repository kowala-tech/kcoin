/* eslint-disable import/no-named-as-default */
import React, { Component } from "react";
import PropTypes from "prop-types";
import { ConnectedRouter } from "react-router-redux";
import { Route } from "react-router-dom";
import { Provider } from "react-redux";
import { IntlProvider } from "react-intl";
import { MuiThemeProvider, withTheme } from "@material-ui/core/styles";
import muiTheme from "./MuiTheme";
import AppContainer from "../containers/AppContainer";
import Particles from "./Particles";

class Root extends Component {
	render() {
		const {
			store,
			history
		} = this.props;

		return (
			<Provider store={store}>
				<IntlProvider locale="en">
					<MuiThemeProvider theme={muiTheme}>
						<Particles />
						<ConnectedRouter history={history}>
							<Route path="/"
								component={AppContainer} />
						</ConnectedRouter>
					</MuiThemeProvider>
				</IntlProvider>
			</Provider>
		);
	}
}

Root.propTypes = {
	store: PropTypes.object.isRequired,
	history: PropTypes.object.isRequired
};

export default withTheme()(Root);
