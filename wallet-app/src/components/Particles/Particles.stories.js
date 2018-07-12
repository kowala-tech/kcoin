import React from "react";
import { storiesOf } from "@storybook/react";
// Story related component imports
import Particles from "./Particles";

const Container = (storyFn) => (
	<div style={{
		backgroundColor: "#1f033d",
		width: "100%",
		height: "100%",
		position: "fixed"
	}}>
		{ storyFn() }
	</div>
);

storiesOf("Particles", module)
	.addDecorator(Container)
	.add("default", () =>
		(<Particles />)
	);
