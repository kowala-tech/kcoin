import { createMuiTheme } from "@material-ui/core/styles";

const muiTheme = createMuiTheme({

	palette: {
		type: "light",
		primary: {
			light: "#8449b5",
			main: "#531c85",
			dark: "#220057",
			contrastText: "#fff"
		},
		secondary: {
			light: "#ff5a80",
			main: "#f90054",
			dark: "#be002c",
			contrastText: "#fff"
		},
		action: {
			hover: "#d4bfe5"
		}
	},
	overrides: {
		MuiButton: {
			root: {
				//borderRadius: 20,
			},
		},
		MuiPaper: {
			rounded: {
				//borderRadius: 10
			}
		}
	}
});

export default muiTheme;
