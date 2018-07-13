import React from "react";
import { storiesOf } from "@storybook/react";
// import { action } from "@storybook/addon-actions";
// Story related component imports
import Snackbar from "./Snackbar";

storiesOf("Snackbar", module)
	.add("default", () =>
		(<Snackbar
			message={"Example message"}
		/>)
	);
