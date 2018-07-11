import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import CircularProgress from "@material-ui/core/CircularProgress";
// Component Related Imports
import styles from "./styles";

class Progress extends React.Component {

	render() {
		const { classes, ...rest } = this.props;

		return (
			<div className={classes.root}>
				<CircularProgress
					{...rest}
					color="secondary"/>
			</div>
		);
	}
}

Progress.propTypes = {
	classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Progress);
