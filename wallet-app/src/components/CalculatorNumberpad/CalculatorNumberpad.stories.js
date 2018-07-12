import React from "react";
import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
// Story related component imports
import CalculatorNumberpad from "./CalculatorNumberpad";

storiesOf("Calculator Numberpad", module)
	.add("default", () =>
		(<CalculatorNumberpad
			handleChange={action()}
		/>)
	);
