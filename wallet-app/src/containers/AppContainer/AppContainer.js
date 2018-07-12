/* eslint-disable import/no-named-as-default */
import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
import { Switch, Route, withRouter, Redirect } from "react-router-dom";
import { connect } from "react-redux";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";
import Header from "../Header";
import LoginContainer from "../LoginContainer";
import SetupContainer from "../SetupContainer";
import WalletContainer from "../WalletContainer";
import ModalContainer from "../ModalContainer";
import MessageContainer from "../MessageContainer";
import { openMessage } from "../MessageContainer/redux/actions";
import { storeEdgeOnWindow, getLocalUsername } from "../../modules/edge";
import { setUsername } from "../../redux/UserRedux";
import PrivateRoute from "../PrivateRoute";

export class App extends React.Component {
	constructor(props) {
		super(props);
		const self = this;
		storeEdgeOnWindow().then( () => {
			getLocalUsername().then( (username) => {
				console.log(username);
				self.props.setUsername(username);
			});
		});
	}

	render() {
		const {
			classes,
			user
		} = this.props;

		return (
			<div className={`kowalaGradient ${classes.root}`}>
				<Header
					user={user}
				/>
				<main className={classes.flexContainer}>
					<Switch>
						<Route path="/"
							exact
							render={() =>
								user.username ? <Redirect to="/login" /> : <SetupContainer />
							}/>
						<Route
							path="/login"
							component={LoginContainer}/>
						<PrivateRoute
							path="/wallet"
							component={WalletContainer}/>
					</Switch>
					<ModalContainer />
					<MessageContainer />
				</main>
				<div className={classes.footer}>
					<Typography
						className={classes.fadedText}
					>
						kWALLET IS OPEN SOURCE SOFTWARE AND IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED.
					</Typography>
				</div>
			</div>
		);
	}
}

App.propTypes = {
	classes: PropTypes.object.isRequired,
	user: PropTypes.object.isRequired,
	setUsername: PropTypes.func.isRequired
};

const mapStateToProps = (state) => {
	return {
		user: state.user
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		setUsername: (username) => {
			if (username) {
				const message = "Welcome back " + username + "!";
				dispatch(setUsername(username));
				dispatch(openMessage(message));
			}
		},
	};
};

export default compose(
	withStyles(styles),
	withRouter,
	connect(mapStateToProps, mapDispatchToProps)
)(App);
