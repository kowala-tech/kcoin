import { createMuiTheme } from "@material-ui/core/styles";

const muiTheme = createMuiTheme({

	palette: {
		type: "light",
		primary: {
			light: "#c5e3f9",
			main: "#008ef5",
			dark: "#008ef5",
			contrastText: "#fff"
		},
		secondary: {
			light: "#33d093",
			main: "#33d093",
			dark: "#33d093",
			contrastText: "#fff"
		},
		action: {
			hover: "#c5e3f9"
		}
	},
	overrides: {
		MuiButton: {
			root: {
				borderRadius: 5,
			},
		},
		MuiPaper: {
			rounded: {
				borderRadius: 5
			}
		}
	}
});

export default muiTheme;
