import React from "react";
import PropTypes from "prop-types";
import PointTarget from "react-point";
import classNames from "classnames";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
// Component Related Imports
import styles from "./styles";

class CalculatorKey extends React.Component {
	render() {
		const {
			children,
			classes,
			className,
			onPress,
			...props
		} = this.props;

		return (
			<PointTarget onPoint={onPress}>
				<Typography
					{...props}
					className={classNames([classes.button, className])}
				>
					{children}
				</Typography>
			</PointTarget>
		);
	}
}

CalculatorKey.propTypes = {
	children: PropTypes.object.isRequired,
	classes: PropTypes.object.isRequired,
	className: PropTypes.string.isRequired,
	onPress: PropTypes.func.isRequired
};

export default withStyles(styles)(CalculatorKey);
