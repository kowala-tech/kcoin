import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import ConfirmationDialog from "./ConfirmationDialog";

describe("ConfirmationDialog", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<ConfirmationDialog
					open={true}
					title={"Discard all changes?"}
					content={"Do you want to discard all changes?"}
					rejectButtonText={"No, keep editing"}
					acceptButtonText={"Yes, discard changes"}
					onAccept={jest.fn()}
					onReject={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
