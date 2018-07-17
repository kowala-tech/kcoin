import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
import moment from "moment";
// Story related component imports
import TransactionList from "./TransactionList";

describe("TransactionList", () => {

	it("it should match the snapshot", () => {
		const transactions = [
			{
				to: "32Be343B94f860124dC4fEe278FDCBD38C102D88",
				from: null,
				nativeAmount: 1000000000000000000,
				fees: 0.01,
				id: "0x6c90256274c87ed74783d1d049304110377005b016af3661c24a71321cf0c0ec",
				timestamp: moment("2018-01-01T00:00:00-00:00").utc()
			}
		];
		const component = toJson(
			shallow(
				<TransactionList
					loading={false}
					transactions={transactions}
					toggleTransactionDetailsDialogFn={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
