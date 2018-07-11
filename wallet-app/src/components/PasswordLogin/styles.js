const styles = (theme) => ({
	flexContainer: {
		display: "flex",
		flexDirection: "column",
		alignItems: "center",
		justifyContent: "center",
		textAlign: "center"
	},
	errorMessage: {
		paddingBottom: theme.spacing.unit * 6,
		height: theme.spacing.unit * 4
	},
	text: {
		fontWeight: 300
	},
});

export default styles;
