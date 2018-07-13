import { usernameAvailable } from "../../../../modules/edge";

const asyncValidate = (values) => {
	return usernameAvailable(values.username).then( (result) => {
		if (!result) {
			throw { username: "Username is already taken." };
		}
	});
};

export default asyncValidate;
