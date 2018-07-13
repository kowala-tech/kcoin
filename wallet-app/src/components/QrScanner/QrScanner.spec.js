import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import CameraIcon from "@material-ui/icons/Camera";
import QrScanner from "./QrScanner";

describe("QrScanner", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<QrScanner
					handleScan={jest.fn()}
					buttonIcon={<CameraIcon />}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
