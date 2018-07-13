import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";
import Modal from "../../components/Modal";
import QrCode from "../../components/QrCode";
import ShareMenu from "../../components/ShareMenu";
import { closeModal } from "../ModalContainer/redux/actions";
import { CopyToClipboard } from "react-copy-to-clipboard";

class DepositModal extends React.Component {

	render() {
		const {
			classes,
			open,
			handleClose,
			depositAddress
		} = this.props;

		const content = (
			<div className={classes.content}>
				<Typography
					align="center"
					variant="subheading"
				>
					Your Wallet's Address
				</Typography>
				<Typography
					align="center"
					variant="body1"
				>
					{depositAddress}
				</Typography>
				<div className={classes.qrCodeContainer}>
					<QrCode text={depositAddress} />
				</div>
				<CopyToClipboard
					text={depositAddress}
					onCopy={() => console.log(`Copied ${depositAddress}!`)}
				>
					<Button
						fullWidth
						variant="raised"
						color="primary"
					>
						Copy Address To Clipboard
					</Button>
				</CopyToClipboard>
				<ShareMenu
					text={depositAddress}
				/>
			</div>
		);

		return (
			<Modal
				open={open}
				title={"Deposit kUSD"}
				content={content}
				handleClose={handleClose}
			/>
		);
	}
}

DepositModal.propTypes = {
	classes: PropTypes.object.isRequired,
	open: PropTypes.bool.isRequired,
	depositAddress: PropTypes.string.isRequired,
	handleClose: PropTypes.func.isRequired
};

const mapStateToProps = (state) => {
	return {
		open: state.ui.modal.open,
		depositAddress: state.wallet.address
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		handleClose: () => {
			dispatch(closeModal());
		}
	};
};

export default compose(
	withStyles(styles),
	connect(mapStateToProps, mapDispatchToProps)
)(DepositModal);
