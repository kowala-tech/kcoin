import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
import classNames from "classnames";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import CircularProgress from "@material-ui/core/CircularProgress";
// Form Imports
import { Field, reduxForm } from "redux-form";
import ReduxFormField from "../../../ReduxFormField";
import validate from "./validate";
// Component Related Imports
import styles from "./styles";

class PasswordForm extends React.Component {
	render() {
		const {
			classes,
			handleSubmit,
			invalid,
			pristine,
			touched,
			error,
			loading
		} = this.props;

		return (
			<form onSubmit={handleSubmit}>
				<Field
					name="username"
					type="username"
					component={ReduxFormField}
					label="Username"
					disabled={loading}
					fullWidth
					className={classNames({
						animated: (touched && error),
						headShake: (touched && error),
					})}
				/>
				<Field
					name="password"
					type="password"
					component={ReduxFormField}
					label="Password"
					disabled={loading}
					fullWidth
					className={classNames({
						animated: (touched && error),
						headShake: (touched && error),
					})}
				/>
				{error && <strong>{error}</strong>}
				<div className={classes.buttonWrapper}>
					<Button
						fullWidth
						variant="raised"
						color="primary"
						type="submit"
						disabled={loading || invalid || pristine}
					>
						{"Login"}
					</Button>
					{loading && <CircularProgress size={24}
						className={classes.buttonProgress} />}
				</div>
			</form>
		);
	}
}

PasswordForm.propTypes = {
	classes: PropTypes.object.isRequired,
	onSubmit: PropTypes.func.isRequired,
	input: PropTypes.object,
	label: PropTypes.object,
	meta: PropTypes.object,
	handleSubmit: PropTypes.func.isRequired,
	invalid: PropTypes.bool,
	pristine: PropTypes.bool,
	error: PropTypes.string,
	touched: PropTypes.bool,
	loading: PropTypes.bool
};

export default compose(
	withStyles(styles),
	reduxForm({
		form: "passwordForm",
		validate
	})
)(PasswordForm);
