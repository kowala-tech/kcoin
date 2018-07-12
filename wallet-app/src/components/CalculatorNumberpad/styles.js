const styles = (theme) => ({
	root: {
		minWidth: "300px",
		borderTop: "1px solid #000",
		paddingTop: theme.spacing.unit * 2
	},
	keypad: {
	},
	inputKeys: {
	},
	digitKeys: {
		display: "flex",
		flex: "1 auto",
		flexFlow: "row wrap-reverse",
	},
	digitKey: {
	},
	operatorKeys: {
		display: "flex",
		flex: "1 auto",
		flexFlow: "row wrap",
	},
	operatorKey: {
		fontSize: 20
	}
});

export default styles;
