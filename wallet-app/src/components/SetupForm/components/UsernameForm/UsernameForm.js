import React from "react";
import { compose } from "recompose";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import CircularProgress from "@material-ui/core/CircularProgress";
// Form Imports
import { Field, reduxForm } from "redux-form";
import asyncValidate from "./asyncValidate";
import validate from "./validate";
import ReduxFormField from "../../../ReduxFormField";
// Component Related Imports
import styles from "../../styles";

class UsernameForm extends React.Component {
	render() {
		const {
			classes,
			handleSubmit,
			asyncValidating,
			invalid,
			pristine
		} = this.props;

		return (
			<form onSubmit={handleSubmit}>
				<Field
					name="username"
					type="text"
					component={ReduxFormField}
					label="Username"
				/>
				<div className={classes.actionsContainer}>
					<Button
						variant="raised"
						color="primary"
						className={classes.button}
						type="submit"
						disabled={invalid || pristine || asyncValidating.length > 0}
					>
						{ asyncValidating ? <CircularProgress size={20}/> : "Next" }
					</Button>
				</div>
			</form>
		);
	}
}

UsernameForm.propTypes = {
	classes: PropTypes.object.isRequired,
	onSubmit: PropTypes.func.isRequired,
	handleSubmit: PropTypes.func.isRequired,
	asyncValidating: PropTypes.oneOfType([
		PropTypes.string,
		PropTypes.bool
	]),
	invalid: PropTypes.bool.isRequired,
	pristine: PropTypes.bool.isRequired,
};

export default compose(
	withStyles(styles),
	reduxForm({
		form: "setupForm",
		destroyOnUnmount: false,
		forceUnregisterOnUnmount: true,
		validate,
		asyncValidate
	})
)(UsernameForm);
