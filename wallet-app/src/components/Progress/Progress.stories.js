import React from "react";
import { storiesOf } from "@storybook/react";
// Story related component imports
import Progress from "./Progress";

storiesOf("Progress", module)
	.add("default", () =>
		<Progress />
	);
