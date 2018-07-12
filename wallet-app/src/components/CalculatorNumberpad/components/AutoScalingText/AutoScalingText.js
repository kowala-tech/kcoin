import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";

class AutoScalingText extends React.Component {
	state = {
		scale: 1
	};

	componentDidUpdate() {
		const { scale } = this.state;

		const node = this.node;
		const parentNode = node.parentNode;

		const availableWidth = parentNode.offsetWidth;
		const actualWidth = node.offsetWidth;
		const actualScale = availableWidth / actualWidth;

		if (scale === actualScale)
			return;

		if (actualScale < 1) {
			this.setScale(actualScale);
		} else if (scale < 1) {
			this.setScale(1);
		}
	}

	setScale = (scale) => {
		this.setState({ scale: scale });
	}

	render() {
		const { scale } = this.state;
		const { classes } = this.props;

		return (
			<div
				className={classes.root}
				style={{ transform: `scale(${scale},${scale})` }}
				ref={node => this.node = node}
			>
				{this.props.children}
			</div>
		);
	}
}

AutoScalingText.propTypes = {
	classes: PropTypes.object.isRequired,
	children: PropTypes.object.isRequired
};


export default withStyles(styles)(AutoScalingText);
