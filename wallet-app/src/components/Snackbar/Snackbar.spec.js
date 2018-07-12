import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import Snackbar from "./Snackbar";

describe("Snackbar", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<Snackbar
					message={"Hello! I am a snackbar."}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
