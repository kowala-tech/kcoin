import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import Progress from "./Progress";

describe("Progress", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<Progress/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
