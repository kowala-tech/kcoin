import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListSubheader from "@material-ui/core/ListSubheader";
import Divider from "@material-ui/core/Divider";
import Typography from "@material-ui/core/Typography";
import Avatar from "@material-ui/core/Avatar";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import Grow from "@material-ui/core/Grow";
import IconButton from "@material-ui/core/IconButton";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import ListIcon from "@material-ui/icons/List";
// Component Related Imports
import ListItem from "./components/ListItem";
import Progress from "../Progress";
import styles from "./styles";
import { unixToHumanDate } from "../../transforms/time";

const groupBy = (array) => {
	return array.reduce(function(groups, item) {
		const timestamp = unixToHumanDate(item["date"]);
		groups[timestamp] = groups[timestamp] || [];
		groups[timestamp].push(item);
		return groups;
	}, {});
};

function TransactionList( { classes, loading, transactions, toggleTransactionDetailsDialogFn, depositButton } ) {
	if (loading) {
		return (<Progress/>);
	}

	const groupedTransactions = groupBy(transactions);

	return (
		<Grow in>
			<Card className={classes.card}>
				<CardHeader
					avatar={
						<Avatar className={classes.avatar}>
							<ListIcon />
						</Avatar>
					}
					action={
						<IconButton disabled>
							<MoreVertIcon />
						</IconButton>
					}
					title={
						<Typography variant="subheading">
							Transactions
						</Typography>
					}
				/>
				{ Object.keys(groupedTransactions).length === 0 && (
					<div>
						<CardContent>
							<Typography variant="body1">It looks like you don't have any kUSD yet! To get started, deposit some kUSD into this wallet.</Typography>
						</CardContent>
						<CardActions
							className={classes.actions}
						>
							{depositButton}
						</CardActions>
					</div>
				)}
				<div>
					<List
						dense
						disablePadding
						className={classes.list}
					>
						{
							Object.keys(groupedTransactions).sort().reverse().map( (date, index) => {
								return(
									<div key={index}>
										<ListSubheader className={classes.stickyHeader}>
											{date}
											<Divider/>
										</ListSubheader>
										{(
											groupedTransactions[date].map( (transaction, index) => {
												return(
													<ListItem
														button
														key={index}
														transaction={transaction}
														onClick={toggleTransactionDetailsDialogFn}
													/>
												);
											})
										)}
									</div>
								);
							})
						}
					</List>
				</div>
			</Card>
		</Grow>
	);
}

TransactionList.propTypes = {
	transactions: PropTypes.array.isRequired,
	loading: PropTypes.bool.isRequired,
	classes: PropTypes.object.isRequired,
	toggleTransactionDetailsDialogFn: PropTypes.func,
	depositButton: PropTypes.object
};

export default withStyles(styles)(TransactionList);
