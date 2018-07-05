import moment from "moment";

export const convertToUnix = (timestamp) => {
	return moment(timestamp).unix();
};

export const unixToHumanDate = (unixTimestamp) => {
	return moment.unix(unixTimestamp).format("LL");
};

export const unixToHumanTime = (unixTimestamp) => {
	return moment.unix(unixTimestamp).format("h:mm:ss A");
};
