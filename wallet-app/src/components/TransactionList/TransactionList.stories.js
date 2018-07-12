import React from "react";
import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";
// Story related component imports
import TransactionList from "./TransactionList";
import transactions from "./testData";

storiesOf("Transaction List", module)
	.add("loading", () =>
		(<TransactionList
			loading={true}
			transactions={[]}
			toggleTransactionDetailsDialogFn={action("transaction-details-click")}
		/>)
	)
	.add("without transactions", () =>
		(<TransactionList
			loading={false}
			transactions={[]}
			toggleTransactionDetailsDialogFn={action("transaction-details-click")}
		/>)
	)
	.add("with transactions", () =>
		(<TransactionList
			loading={false}
			transactions={transactions}
			toggleTransactionDetailsDialogFn={action("transaction-details-click")}
		/>)
	);
