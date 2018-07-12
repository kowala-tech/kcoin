const drawerWidth = "75vw";
const maxDrawerWidth = "300px";
const toolbarHeight = 75;

const styles = theme => ({
	appFrame: {
		position: "relative",
		width: "100%",
		maxWidth: "720px",
		height: "100%",
		marginTop: toolbarHeight,
		marginBottom: theme.spacing.unit * 2,
		overflow: "scroll"
	},
	drawerPaper: {
		width: drawerWidth,
		maxWidth: maxDrawerWidth,
		height: "100%",
		[theme.breakpoints.up("md")]: {
			maxWidth: maxDrawerWidth,
			height: "auto",
			marginRight: theme.spacing.unit * 2
		}
	},
	content: {
		width: "100%",
		marginBottom: theme.spacing.unit * 2
	},
	logo: {
		height: toolbarHeight - (theme.spacing.unit * 4),
	},
	avatarIcon: {
		height: "45px",
		width: "45px"
	},
});

export default styles;
