import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";
import AutoScalingText from "../AutoScalingText";

class CalculatorDisplay extends React.Component {
	render() {
		const {
			classes,
			value,
			...props
		} = this.props;

		const language = navigator.language || "en-US";
		let formattedValue = parseFloat(value).toLocaleString(language, {
			useGrouping: true,
			maximumFractionDigits: 6
		});

		// Add back missing .0 in e.g. 12.0
		const match = value.match(/\.\d*?(0*)$/);

		if (match)
			formattedValue += (/[1-9]/).test(match[0]) ? match[1] : match[0];

		return (
			<div
				{...props}
				className={classes.root}
			>
				<AutoScalingText>{formattedValue}</AutoScalingText>
			</div>
		);
	}
}

CalculatorDisplay.propTypes = {
	classes: PropTypes.object.isRequired,
	value: PropTypes.object.isRequired
};


export default withStyles(styles)(CalculatorDisplay);
