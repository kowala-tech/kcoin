import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Zoom from "@material-ui/core/Zoom";
// Component Related Imports
import styles from "./styles";

class Fab extends React.Component {
	render() {
		const {
			classes,
			theme,
			icon,
			onClick,
			color
		} = this.props;

		const transitionDuration = {
			enter: theme.transitions.duration.enteringScreen,
			exit: theme.transitions.duration.leavingScreen,
		};

		return (
			<Zoom
				appear={true}
				in={true}
				timeout={transitionDuration}
				style={{
					transitionDelay: transitionDuration.exit
				}}
				unmountOnExit
			>
				<Button
					variant="fab"
					className={classes.fab}
					color={color}
					onClick={onClick}
				>
					{icon}
				</Button>
			</Zoom>
		);
	}
}

Fab.propTypes = {
	classes: PropTypes.object.isRequired,
	theme: PropTypes.object.isRequired,
	onClick: PropTypes.func.isRequired,
	icon: PropTypes.element.isRequired,
	color: PropTypes.string
};

export default withStyles(styles, { withTheme: true })(Fab);
