import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import Button from "@material-ui/core/Button";
// Component Related Imports
import Modal from "../Modal";

class ConfirmationDialog extends React.Component {
	render() {

		const {
			open,
			title,
			content,
			rejectButtonText,
			acceptButtonText,
			onReject,
			onAccept,
		} = this.props;

		return (
			<Modal
				open={open}
				handleClose={onReject}
				title={title}
				content={content}
				rejectButton={(
					<Button onClick={onReject}>
						{rejectButtonText}
					</Button>
				)}
				acceptButton={(
					<Button onClick={onAccept}
						autoFocus
						color="secondary"
						variant="raised">
						{acceptButtonText}
					</Button>
				)}
			/>
		);
	}
}

ConfirmationDialog.propTypes = {
	open: PropTypes.bool.isRequired,
	title: PropTypes.string.isRequired,
	content: PropTypes.string.isRequired,
	rejectButtonText: PropTypes.string.isRequired,
	acceptButtonText: PropTypes.string.isRequired,
	onAccept: PropTypes.func.isRequired,
	onReject: PropTypes.func.isRequired
};

export default compose(
)(ConfirmationDialog);
