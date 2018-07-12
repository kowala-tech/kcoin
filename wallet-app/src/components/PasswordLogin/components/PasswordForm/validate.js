const validate = values => {
	const errors = {};
	// Username
	if (!values.username) {
		errors.username = "Username is required.";
	}
	// Password
	if (!values.password) {
		errors.password = "Password is required.";
	}
	return errors;
};

export default validate;
