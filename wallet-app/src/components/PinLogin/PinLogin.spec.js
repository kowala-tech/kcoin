import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import PinLogin from "./PinLogin";

describe("PinLogin", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<PinLogin
					handleSubmit={jest.fn()}
					loading={false}
					error={false}
					errorMessage={""}
					username={"test-user"}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
