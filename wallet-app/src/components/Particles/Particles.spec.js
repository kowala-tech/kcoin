import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import Particles from "./Particles";

describe("Particles", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<Particles/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
