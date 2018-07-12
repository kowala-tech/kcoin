import React from "react";
import { storiesOf } from "@storybook/react";
// Story related component imports
import BalanceCard from "./BalanceCard";

const ContrastBackground = (storyFn) => (
	<div style={{ backgroundColor: "#1f033d" }}>
		{ storyFn() }
	</div>
);

storiesOf("Balance Card", module)
	.addDecorator(ContrastBackground)
	.add("with balance", () => (
		<BalanceCard
			balance={12.3456789}
		/>
	));
