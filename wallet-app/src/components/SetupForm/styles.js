const styles = (theme) => ({
	root: {
		width: "100%",
		maxWidth: "500px"
	},
	button: {
		marginTop: theme.spacing.unit,
		marginRight: theme.spacing.unit,
	},
	stepper: {
		paddingLeft: 0,
		paddingRight: 0
	},
	actionsContainer: {
		textAlign: "right",
		marginBottom: theme.spacing.unit * 2,
	},
	pendingContainer: {
		display: "flex",
		alignItems: "center",
	},
	progress: {
		marginRight: theme.spacing.unit * 2,
		minWidth: 24
	},
});

export default styles;
