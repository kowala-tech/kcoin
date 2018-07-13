import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import BalanceCard from "./BalanceCard";

describe("BalanceCard", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<BalanceCard
					balance={"43.23438"}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
