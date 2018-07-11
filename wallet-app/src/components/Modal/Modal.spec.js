import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import Modal from "./Modal";
import Button from "@material-ui/core/Button";

describe("Modal", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<Modal
					open={true}
					title={"Modal Title"}
					content={"Modal content."}
					rejectButton={(
						<Button
							onClick={jest.fn()}
						>
							No
						</Button>
					)}
					acceptButton={(
						<Button
							onClick={jest.fn()}
						>
							Yes
						</Button>
					)}
					handleClose={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
