import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
import classNames from "classnames";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import LockOutline from "@material-ui/icons/LockOutline";
import Grow from "@material-ui/core/Grow";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";
import PinForm from "./components/PinForm";

class PinLogin extends React.Component {

	render() {
		const {
			classes,
			handleSubmit,
			error,
			//errorMessage,
			loading,
			username
		} = this.props;

		return (
			<div>
				<Grow in>
					<div className={classes.flexContainer}>
						<LockOutline className={classNames("kowala",[classes.lockIcon])}/>
						<PinForm
							onSubmit={handleSubmit}
							loading={loading}
							error={error}
							username={username}
						/>
						<Typography variant="body1"
							className={classes.unlockMessage}>
							Enter your PIN to unlock your wallet
						</Typography>
					</div>
				</Grow>
			</div>
		);
	}
}

PinLogin.propTypes = {
	classes: PropTypes.object.isRequired,
	handleSubmit: PropTypes.func.isRequired,
	loading: PropTypes.bool.isRequired,
	error: PropTypes.bool.isRequired,
	errorMessage: PropTypes.string,
	username: PropTypes.string.isRequired
};

export default compose(
	withStyles(styles),
)(PinLogin);
