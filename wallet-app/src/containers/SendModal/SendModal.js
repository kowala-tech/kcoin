import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
import { Field, reduxForm, change } from "redux-form";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Avatar from "@material-ui/core/Avatar";
import FormControl from "@material-ui/core/FormControl";
import InputAdornment from "@material-ui/core/InputAdornment";
import InputLabel from "@material-ui/core/InputLabel";
import CloseIcon from "@material-ui/icons/Close";
import IconButton from "@material-ui/core/IconButton";
import Input from "@material-ui/core/Input";
// Component Related Imports
import styles from "./styles";
import { sendTransaction } from "../../redux/WalletRedux/";
import Modal from "../../components/Modal";
// import CalculatorNumberpad from "../../components/CalculatorNumberpad";
import { closeModal } from "../ModalContainer/redux/actions";
import { collapseAddress } from "../../transforms/address";
import { kusdToWei } from "../../transforms/currency";

class SendModal extends React.Component {
	render() {

		const AddressField = ({
			input,
			meta: { dirty },
			...custom
		}) => (
			<FormControl fullWidth>
				<InputLabel>Where do you want to send kUSD?</InputLabel>
				<Input
					type="text"
					placeholder={"Paste or scan a wallet address"}
					endAdornment={
						<InputAdornment position="end">
							{ dirty && <IconButton
								onClick={this.props.reset}
							>
								<CloseIcon />
							</IconButton> }
						</InputAdornment>
					}
					{...input}
					{...custom}
				/>
			</FormControl>
		);

		const AmountField = ({
			input,
			meta: { dirty },
			...custom
		}) => (
			<FormControl fullWidth>
				<InputLabel>How much kUSD do you want to send?</InputLabel>
				<Input
					type="number"
					placeholder={"Enter an amount"}
					endAdornment={
						<InputAdornment position="end">
							{ dirty && <IconButton
								onClick={this.props.reset}
							>
								<CloseIcon />
							</IconButton> }
						</InputAdornment>
					}
					{...input}
					{...custom}
				/>
			</FormControl>
		);

		const {
			classes,
			open,
			handleClose,
			handleSubmit,
			pristine,
			submitting,
			setAddress
		} = this.props;

		const button = (
			<Button
				type="sumbit"
				fullWidth
				size="large"
				variant="raised"
				color="primary"
				disabled={submitting || pristine}
			>
				Continue
			</Button>
		);

		const content = (
			<div className={classes.root}>
				<form
					onSubmit={handleSubmit}
					className={classes.form}
					autoComplete="off"
				>
					<div>
						<Field
							name="address"
							component={AddressField}
						/>
						<List>
							<ListItem
								dense
								disableGutters
								button
								onClick={() => setAddress("0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe")}
							>
								<Avatar>0x</Avatar>
								<ListItemText
									primary="Anonymous User"
									secondary={collapseAddress("0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe")} />
							</ListItem>
						</List>
					</div>
					<div>
						<Field
							name="amount"
							component={AmountField}
						/>
					</div>
					{button}
				</form>
			</div>
		);

		return (
			<Modal
				open={open}
				title={"Send kUSD"}
				content={content}
				handleClose={handleClose}
			/>
		);
	}
}

SendModal.propTypes = {
	classes: PropTypes.object.isRequired,
	open: PropTypes.bool.isRequired,
	handleClose: PropTypes.func.isRequired,
	onSubmit: PropTypes.func.isRequired,
	walletId: PropTypes.string.isRequired,
	amount: PropTypes.number,
	address: PropTypes.string,
	change: PropTypes.func.isRequired,
	handleSubmit: PropTypes.func.isRequired,
	initialValues: PropTypes.object.isRequired,
	reset: PropTypes.func.isRequired,
	setAddress: PropTypes.func.isRequired,
	pristine: PropTypes.bool.isRequired,
	submitting: PropTypes.bool.isRequired
};

const mapStateToProps = (state) => {
	return {
		open: state.ui.modal.open,
		walletId: state.wallet.id,
		sendingWalletAddress: state.wallet.address,
		initialValues: { walletId: state.wallet.id }
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		handleClose: () => {
			dispatch(closeModal());
		},
		setAddress: (address) => {
			if (address) {
				dispatch(change("send", "address", address));
			}
		},
		onSubmit: (values) => {
			dispatch(sendTransaction(
				values.walletId,
				values.address,
				kusdToWei(values.amount)
			));
			dispatch(closeModal());
		}
	};
};

export default compose(
	withStyles(styles),
	connect(mapStateToProps, mapDispatchToProps),
	reduxForm({
		form: "send"
	})
)(SendModal);
