import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import Fab from "./Fab";
import HomeIcon from "@material-ui/icons/Home";

describe("Fab", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<Fab
					icon={<HomeIcon />}
					onClick={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
