const validate = values => {
	const errors = {};
	// Username
	if (!values.username) {
		errors.username = "Username is required.";
	} else if (values.username.length < 8) {
		errors.username = "Username must be at least 8 characters.";
	}
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
	// Pin
	if (!values.pin) {
		errors.pin = "PIN is required.";
	} else if (values.pin.length !== 4) {
		errors.pin = "PIN must 4 digits.";
	}
	// Pin Confirmation
	if (values.pin !== values.pinConfirmation) {
		errors.pinConfirmation = "PINs do not match.";
	}
	return errors;
};

export default validate;
