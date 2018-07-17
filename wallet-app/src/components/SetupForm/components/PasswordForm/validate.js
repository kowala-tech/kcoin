const validate = values => {
	const errors = {};
	// Password
	if (!values.password) {
		errors.password = "Password is required.";
	} else if (values.password.length < 8) {
		errors.password = "Password must be at least 8 characters.";
	}
	// Password Confirmation
	if (values.password !== values.passwordConfirmation) {
		errors.passwordConfirmation = "Passwords do not match.";
	}
	return errors;
};

export default validate;
