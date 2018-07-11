import React from "react";
import PropTypes from "prop-types";
import { FormattedNumber } from "react-intl";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";
import { weiToKusd } from "../../transforms/currency";

class BalanceCard extends React.Component {

	render() {
		const {
			classes,
			balance
		} = this.props;

		return (
			<div className={classes.root}>
				<Typography
					className={classes.text}
				>
					Your wallet's current balance is
				</Typography>
				<Typography
					variant="title"
					className={classes.balance}
				>
					<FormattedNumber
						minimumFractionDigits={4}
						maximumFractionDigits={18}
						value={weiToKusd(balance)}
					/>
					{" kUSD"}
				</Typography>
			</div>
		);
	}
}

BalanceCard.propTypes = {
	classes: PropTypes.object.isRequired,
	balance: PropTypes.string
};

export default withStyles(styles)(BalanceCard);
