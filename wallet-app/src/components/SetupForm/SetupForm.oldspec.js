import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import SetupForm from "./SetupForm";

describe("SetupForm", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<SetupForm
					onSubmit={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
