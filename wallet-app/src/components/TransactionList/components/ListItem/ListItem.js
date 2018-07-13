import React from "react";
import PropTypes from "prop-types";
import { FormattedNumber } from "react-intl";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import ListItem from "@material-ui/core/ListItem";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import ListItemText from "@material-ui/core/ListItemText";
// Icons
import Avatar from "@material-ui/core/Avatar";
import PersonIcon from "@material-ui/icons/Person";
// Component Related Imports
import styles from "./styles";
import { weiToKusd } from "../../../../transforms/currency";
import { unixToHumanTime } from "../../../../transforms/time";

const TransactionListItem = ({
	classes,
	transaction,
	toggleTransactionDetailsDialogFn
}) => {
	return (
		<ListItem button
			onClick={toggleTransactionDetailsDialogFn}>
			<ListItemAvatar>
				<Avatar className={classes.avatar}>
					<PersonIcon />
				</Avatar>
			</ListItemAvatar>
			<ListItemText
				className={classes.truncate}
				primary={transaction.txid.substring(0,16)}
				secondary={unixToHumanTime(transaction.date)}
			/>
			<ListItemText
				className={classes.amount}
				primary={
					<FormattedNumber
						minimumFractionDigits={4}
						maximumFractionDigits={18}
						value={weiToKusd(transaction.nativeAmount)}
					/>
				}
				secondary="kUSD"
			/>
		</ListItem>
	);
};

TransactionListItem.propTypes = {
	classes: PropTypes.object.isRequired,
	transaction: PropTypes.object,
	toggleTransactionDetailsDialogFn: PropTypes.func,
};

export default withStyles(styles)(TransactionListItem);
