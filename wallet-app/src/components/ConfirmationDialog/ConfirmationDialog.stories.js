import React from "react";
import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
// Story related component imports
import ConfirmationDialog from "./ConfirmationDialog";

storiesOf("Confirmation Dialog", module)
	.add("with example content", () =>
		(<ConfirmationDialog
			open={true}
			title={"Discard all changes?"}
			content={"Do you want to discard all changes?"}
			rejectButtonText={"No, keep editing"}
			acceptButtonText={"Yes, discard changes"}
			onAccept={action("onAccept")}
			onReject={action("onReject")}
		/>)
	);
