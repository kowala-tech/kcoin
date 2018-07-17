import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import withMobileDialog from "@material-ui/core/withMobileDialog";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
// Component Related Imports
import styles from "./styles";

class Modal extends React.Component {
	render() {

		const {
			open,
			handleClose,
			title,
			content,
			acceptButton,
			rejectButton,
			fullScreen,
			classes
		} = this.props;

		return (
			<Dialog
				open={open}
				onClose={handleClose}
				onBackdropClick={handleClose}
				fullScreen={fullScreen}
			>
				<AppBar position="static">
					<Toolbar>
						<Typography
							variant="title"
							color="inherit"
							className={classes.flex}
						>
							{title}
						</Typography>
						<IconButton
							onClick={handleClose}
							className={classes.menuButton}
							color="inherit"
						>
							<CloseIcon />
						</IconButton>
					</Toolbar>
				</AppBar>
				<DialogContent classes={{ root: classes.content }}>
					{content}
				</DialogContent>
				<DialogActions>
					{rejectButton}
					{acceptButton}
				</DialogActions>
			</Dialog>
		);
	}
}

Modal.propTypes = {
	classes: PropTypes.object.isRequired,
	open: PropTypes.bool.isRequired,
	handleClose: PropTypes.func.isRequired,
	title: PropTypes.string.isRequired,
	content: PropTypes.object.isRequired,
	acceptButton: PropTypes.object,
	rejectButton: PropTypes.object,
	fullScreen: PropTypes.bool.isRequired
};

export default compose(
	withStyles(styles),
	withMobileDialog()
)(Modal);
