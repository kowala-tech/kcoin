import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import SendIcon from "@material-ui/icons/Send";
import Button from "@material-ui/core/Button";
// Component Related Imports
import styles from "./styles";
import Fab from "../../components/Fab";
import TransactionList from "../../components/TransactionList";
import BalanceCard from "../../components/BalanceCard";
import { openModal } from "../ModalContainer/redux/actions";

class WalletContainer extends React.Component {

	render() {
		const {
			classes,
			openModal,
			wallet
		} = this.props;

		return (
			<div className={classes.appFrame}>
				<main className={classes.content}>
					<BalanceCard balance={wallet.balance}/>
					<TransactionList
						loading={wallet.loading}
						transactions={wallet.transactions}
						depositButton={(
							<Button
								fullWidth
								variant="raised"
								color="primary"
								onClick={() => {openModal("DEPOSIT_MODAL");}}
							>
								Deposit kUSD
							</Button>
						)}
					/>
				</main>
				{ wallet.balance > 0 && (
					<Fab
						color="primary"
						onClick={() => {openModal("SEND_MODAL");}}
						icon={<SendIcon />}
					/>
				) }
			</div>
		);
	}
}

WalletContainer.propTypes = {
	classes: PropTypes.object.isRequired,
	openModal: PropTypes.func.isRequired,
	wallet: PropTypes.object.isRequired
};

const mapStateToProps = (state) => {
	return {
		wallet: state.wallet
	};
};

const mapDispatchToProps = (dispatch) => {
	return ({
		openModal: (modelType) => {
			dispatch(openModal(modelType));
		}
	});
};

export default compose(
	withStyles(styles),
	connect(mapStateToProps, mapDispatchToProps)
)(WalletContainer);
