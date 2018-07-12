import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import Grow from "@material-ui/core/Grow";
// Component Related Imports
import styles from "./styles";
import PinLogin from "../../components/PinLogin";
import PasswordLogin from "../../components/PasswordLogin";
import { loginWithPin, loginWithPassword } from "../../redux/UserRedux";

class LoginContainer extends React.Component {
	render() {
		const {
			loginWithPin,
			loginWithPassword,
			error,
			loading,
			errorMessage,
			returningUser,
			username
		} = this.props;

		const loginParams = {
			error,
			loading,
			errorMessage
		};

		return (
			<div>
				{ returningUser ?
					<PinLogin
						{...loginParams}
						handleSubmit={loginWithPin}
						username={username}
					/>
					:
					<Grow in>
						<Card>
							<CardContent>
								<Typography variant="subheading">Login To Your Account</Typography>
								<PasswordLogin
									{...loginParams}
									handleSubmit={loginWithPassword}
								/>
							</CardContent>
						</Card>
					</Grow>
				}
			</div>
		);
	}
}

LoginContainer.propTypes = {
	loginWithPin: PropTypes.func.isRequired,
	loginWithPassword: PropTypes.func.isRequired,
	loading: PropTypes.bool.isRequired,
	error: PropTypes.bool.isRequired,
	errorMessage: PropTypes.string,
	returningUser: PropTypes.bool,
	username: PropTypes.string
};

const mapStateToProps = (state) => {
	return {
		loading: state.user.loading,
		error: state.user.error,
		errorMessage: state.user.errorMessage,
		returningUser: state.user.username ? state.user.username.length > 0 : false,
		username: state.user.username
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		loginWithPin: (username, pin, callback) => {
			dispatch(loginWithPin(
				username,
				pin,
				callback
			));
		},
		loginWithPassword: (props) => {
			dispatch(loginWithPassword(
				props.username,
				props.password
			));
		}
	};
};

export default compose(
	withStyles(styles),
	connect(mapStateToProps, mapDispatchToProps)
)(LoginContainer);
