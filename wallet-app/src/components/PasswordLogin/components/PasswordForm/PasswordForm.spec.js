import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import PasswordForm from "./PasswordForm";

describe("PasswordForm", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<PasswordForm
					onSubmit={jest.fn()}
					handleSubmit={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
