const toolbarHeight = 75;

const styles = (theme) => ({
	root: {
		height: "inherit",
		backgroundColor: theme.palette.primary.main,
		color: theme.palette.primary.contrastText
	},
	flexContainer: {
		display: "flex",
		flexDirection: "column",
		alignItems: "center",
		justifyContent: "center",
		height: "100vh",
		marginLeft: theme.spacing.unit * 2,
		marginRight: theme.spacing.unit * 2,
		marginBottom: theme.spacing.unit * 2,
		marginTop: -toolbarHeight
	},
	footer: {
		position: "fixed",
		bottom: 0,
		width: `calc(100% - ${theme.spacing.unit*4}px)`,
		margin: theme.spacing.unit * 2,
		textAlign: "center"
	},
	fadedText: {
		color: theme.palette.primary.contrastText,
		opacity: "0.75",
		fontSize: 10
	}
});

export default styles;
