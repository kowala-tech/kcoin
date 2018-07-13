export const closeMessage = () => {
	return { type: "CLOSE_MESSAGE" };
};

export const openMessage = (text) => {
	return {
		type: "OPEN_MESSAGE",
		text
	};
};
