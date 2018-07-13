const validate = values => {
	const errors = {};
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
