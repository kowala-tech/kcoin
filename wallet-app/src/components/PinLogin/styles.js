// const toolbarHeight = 75;

const styles = (theme) => ({
	flexContainer: {
		display: "flex",
		flexDirection: "column",
		alignItems: "center",
		justifyContent: "center",
		textAlign: "center"
	},
	lockIcon: {
		marginBottom: theme.spacing.unit * 6
	},
	unlockMessage: {
		marginTop: theme.spacing.unit * 6,
		textAlign: "center",
		color: theme.palette.primary.contrastText,
		fontWeight: 200
	},
	close: {
		width: theme.spacing.unit * 4,
		height: theme.spacing.unit * 4,
	},
});

export default styles;
