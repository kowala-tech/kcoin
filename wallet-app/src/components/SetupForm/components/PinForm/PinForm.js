import React from "react";
import { compose } from "recompose";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
// Form Imports
import { Field, reduxForm } from "redux-form";
import validate from "./validate";
import ReduxFormField from "../../../ReduxFormField";
// Component Related Imports
import styles from "../../styles";

class PinForm extends React.Component {
	render() {
		const { classes, handleSubmit, invalid, pristine, previousStep } = this.props;
		return (
			<form onSubmit={handleSubmit}>
				<Field
					name="pin"
					type="password"
					component={ReduxFormField}
					label="PIN"
				/>
				<Field
					name="pinConfirmation"
					type="password"
					component={ReduxFormField}
					label="PIN Confirmation"
				/>
				<div className={classes.actionsContainer}>
					<Button
						onClick={previousStep}
						className={classes.button}
					>
						Back
					</Button>
					<Button
						variant="raised"
						color="primary"
						className={classes.button}
						type="submit"
						disabled={invalid || pristine}
					>
						Finish
					</Button>
				</div>
			</form>
		);
	}
}

PinForm.propTypes = {
	classes: PropTypes.object.isRequired,
	onSubmit: PropTypes.func.isRequired,
	previousStep: PropTypes.func.isRequired,
	handleSubmit: PropTypes.func.isRequired,
	invalid: PropTypes.bool.isRequired,
	pristine: PropTypes.bool.isRequired,
};

export default compose(
	withStyles(styles),
	reduxForm({
		form: "setupForm",
		destroyOnUnmount: false,
		forceUnregisterOnUnmount: true,
		validate
	})
)(PinForm);
