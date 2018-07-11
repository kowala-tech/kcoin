import React from "react";
import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
// Story related component imports
import Fab from "./Fab";
import HomeIcon from "@material-ui/icons/Home";

storiesOf("FAB", module)
	.add("primary color", () =>
		(<Fab
			icon={<HomeIcon />}
			onClick={action("fab-click")}
			color="primary"
		/>)
	)
	.add("secondary color", () =>
		(<Fab
			icon={<HomeIcon />}
			onClick={action("fab-click")}
			color="secondary"
		/>)
	);
