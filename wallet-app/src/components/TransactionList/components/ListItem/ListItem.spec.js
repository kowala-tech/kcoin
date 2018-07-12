import React from "react";
import { shallow } from "enzyme";
import toJson from "enzyme-to-json";
// Story related component imports
import moment from "moment";
import ListItem from "./ListItem";

describe("ListItem", () => {

	it("it should match the snapshot", () => {
		const component = toJson(
			shallow(
				<ListItem
					amount={12.8544321}
					timestamp={moment("2018-01-01T00:00:00-00:00").utc().unix()}
					index={0}
					toggleTransactionDetailsDialogFn={jest.fn()}
				/>
			)
		);
		expect(component).toMatchSnapshot();
	});

});
