import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";
import PasswordForm from "./components/PasswordForm";

class PasswordLogin extends React.Component {
	render() {
		const {
			classes,
			handleSubmit,
			loading,
			errorMessage
		} = this.props;

		return (
			<div className={classes.flexContainer}>
				<Typography className={classes.text}>{errorMessage}</Typography>
				<PasswordForm
					loading={loading}
					onSubmit={handleSubmit}
				/>
			</div>
		);
	}

}

PasswordLogin.propTypes = {
	classes: PropTypes.object.isRequired,
	handleSubmit: PropTypes.func.isRequired,
	loading: PropTypes.bool.isRequired,
	error: PropTypes.bool.isRequired,
	errorMessage: PropTypes.string
};

export default compose(
	withStyles(styles)
)(PasswordLogin);
