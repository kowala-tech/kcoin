import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import QrCode from "./QrCode";

describe("QrCode", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<QrCode
					text="this is text"
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
