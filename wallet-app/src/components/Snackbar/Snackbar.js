import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import MuiSnackbar from "@material-ui/core/Snackbar";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
// Component Related Imports
import styles from "./styles";

class Snackbar extends React.Component {
	render() {
		const {
			classes,
			open,
			handleClose,
			message
		} = this.props;

		return (
			<MuiSnackbar
				anchorOrigin={{
					vertical: "bottom",
					horizontal: "right",
				}}
				open={open}
				autoHideDuration={5000}
				onClose={handleClose}
				onExited={handleClose}
				message={<span id="message-id">{message}</span>}
				action={[
					<IconButton
						key="close"
						aria-label="Close"
						color="inherit"
						className={classes.close}
						onClick={handleClose}
					>
						<CloseIcon />
					</IconButton>
				]}
			/>
		);
	}
}

Snackbar.propTypes = {
	classes: PropTypes.object.isRequired,
	message: PropTypes.string,
	open: PropTypes.bool,
	handleClose: PropTypes.func
};

export default withStyles(styles)(Snackbar);
