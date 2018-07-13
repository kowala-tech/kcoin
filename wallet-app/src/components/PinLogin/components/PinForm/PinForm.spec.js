import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import PinForm from "./PinForm";

describe("PinForm", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<PinForm
					onSubmit={jest.fn()}
					loading={false}
					error={false}
					username={"test-user"}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
