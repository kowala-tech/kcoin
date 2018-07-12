import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";
import { QRCode } from "react-qr-svg";

class QrCode extends React.Component {

	render() {
		const {
			classes,
			text,
			...rest
		} = this.props;

		return (
			<div className={classes.root}>
				<QRCode
					value={text}
					{...rest}
				/>
			</div>
		);
	}
}

QrCode.propTypes = {
	classes: PropTypes.object.isRequired,
	text: PropTypes.string.isRequired
};

export default withStyles(styles)(QrCode);
