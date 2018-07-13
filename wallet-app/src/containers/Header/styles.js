const styles = (theme) => ({
	appBar: {
		display: "flex",
		flexGrow: 1
	},
	logo: {
		height: "24px",
		margin: `${theme.spacing.unit} 0`
	},
	menuLogo: {
		height: "18px",
		marginTop: theme.spacing.unit
	},
	header: {
		marginBottom: theme.spacing.unit
	},
	avatar: {
		marginRight: theme.spacing.unit * 2,
		textTransform: "uppercase"
	},
	icon: {
		color: theme.palette.primary.contrastText,
		opacity: "0.9"
	},
	leftIcon: {
		flex: 1
	},
	title: {
		flex: 2,
		textAlign: "center"
	},
	rightIcon: {
		flex: 1,
		textAlign: "right"
	}
});

export default styles;
