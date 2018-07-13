import moment from "moment";

const buildTransactions = (n=30) => {
	let arr = [];
	for (let i = 0; i < n; i++) {
		arr.push(
			{
				to: "32Be343B94f860124dC4fEe278FDCBD38C102D88",
				from: null,
				nativeAmount: 1000000000000000000,
				fees: 0.01,
				id: "0x6c90256274c87ed74783d1d049304110377005b016af3661c24a71321cf0c0ec",
				timestamp: moment().subtract(Math.floor(Math.random() * Math.floor(14)), "day").format()
			}
		);
	}
	return arr;
};

const transactions = buildTransactions();

export default transactions;
