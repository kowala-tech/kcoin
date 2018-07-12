import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import PasswordLogin from "./PasswordLogin";

describe("PasswordLogin", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<PasswordLogin
					handleSubmit={jest.fn()}
					loading={false}
					error={false}
					errorMessage={""}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
