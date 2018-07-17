const styles = (theme) => ({
	card: {
		marginBottom: theme.spacing.unit * 2,
		paddingBottom: theme.spacing.unit
	},
	stickyHeader: {
		backgroundColor: theme.palette.background.paper,
	},
	list: {
		width: "100%",
		backgroundColor: theme.palette.background.paper,
		position: "relative",
		overflow: "auto"
	},
	avatar: {
		backgroundColor: theme.palette.primary.main
	}
});

export default styles;
