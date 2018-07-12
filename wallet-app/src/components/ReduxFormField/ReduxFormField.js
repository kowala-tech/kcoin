import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import TextField from "@material-ui/core/TextField";

const ReduxFormField = ({
	input,
	label,
	meta: { touched, error },
	...custom
}) => (
	<TextField
		fullWidth
		label={label}
		error={touched && error && error.length > 0}
		helperText={touched && error}
		FormHelperTextProps={{ error: error && error.length > 0 }}
		{...input}
		{...custom}
	/>
);

ReduxFormField.propTypes = {
	input: PropTypes.object.isRequired,
	label: PropTypes.string.isRequired,
	meta: PropTypes.object.isRequired,
	custom: PropTypes.object
};

export default ReduxFormField;
