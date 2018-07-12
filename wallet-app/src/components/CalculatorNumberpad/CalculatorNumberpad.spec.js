import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import CalculatorNumberpad from "./CalculatorNumberpad";

describe("CalculatorNumberpad", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<CalculatorNumberpad
					handleChange={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
