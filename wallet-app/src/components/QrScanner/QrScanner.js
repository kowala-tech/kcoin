import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import Button from "@material-ui/core/Button";
// Component Related Imports
import styles from "./styles";
import QrReader from "react-qr-reader";

class QrScanner extends React.Component {
	constructor(props){
		super(props);
		this.state = {
			delay: 300,
			open: false
		};
	}

	handleError(err){
		console.error(err);
	}

	toggle(){
		this.setState({ open: !this.state.open });
	}


	render(){

		const {
			handleScan
		} = this.props;

		const scanAndClose = (result) => {
			if (result && result.length > 10) {
				handleScan(result) && this.toggle;
			}
		};

		return(
			<div>
				<Button
					variant="raised"
					fullWidth
					color="primary"
					onClick={this.toggle.bind(this)}
				>
					Scan Wallet Address QR Code
				</Button>
				<Drawer
					anchor="bottom"
					open={this.state.open}
					onClose={this.toggle.bind(this)}
				>
					<div
						role="button"
						onClick={this.toggle.bind(this)}
						onKeyDown={this.toggle.bind(this)}
					>
						<QrReader
							showViewFinder={false}
							delay={this.state.delay}
							onError={this.handleError}
							onScan={(result) => scanAndClose(result)}
							style={{ width: "100%" }}
						/>
					</div>
				</Drawer>
			</div>
		);
	}
}

QrScanner.propTypes = {
	classes: PropTypes.object.isRequired,
	handleScan: PropTypes.func.isRequired
};

export default compose(
	withStyles(styles),
)(QrScanner);
