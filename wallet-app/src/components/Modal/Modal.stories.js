import React from "react";
import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
// Story related component imports
import Modal from "./Modal";
import Button from "@material-ui/core/Button";

storiesOf("Modal", module)
	.add("with example content", () =>
		(<Modal
			open={true}
			title={"Example Modal"}
			content={"This is an example modal used for displaying content over an the current page."}
			rejectButton={(
				<Button
					onClick={action("onReject")}
				>
					No means no!
				</Button>
			)}
			acceptButton={(
				<Button
					variant="raised"
					color="secondary"
					onClick={action("onAccept")}
				>
					Thanks for the rad info brosef!
				</Button>
			)}
			handleClose={action("onReject")}
		/>)
	);
