import React from "react";
import { storiesOf } from "@storybook/react";
// Story related component imports
import QrCode from "./QrCode";

storiesOf("QrCode", module)
	.add("default", () =>
		(<QrCode
			text="this is text"
		/>)
	);
