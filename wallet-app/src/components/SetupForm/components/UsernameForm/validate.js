const validate = values => {
	const errors = {};
	// Username
	if (!values.username) {
		errors.username = "Username is required.";
	} else if (values.username.length < 8) {
		errors.username = "Username must be at least 8 characters.";
	}
	return errors;
};

export default validate;
