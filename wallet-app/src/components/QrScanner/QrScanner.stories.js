import React from "react";
import { storiesOf } from "@storybook/react";
// import { action } from "@storybook/addon-actions";
// Story related component imports
import QrScanner from "./QrScanner";

storiesOf("QR Scanner", module)
	.add("default", () =>
		(<QrScanner/>)
	);
