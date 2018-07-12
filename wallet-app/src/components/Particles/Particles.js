import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";
import ParticlesJS from "react-particles-js";
import particlesConfig from "./particlesConfig";

class Particles extends React.Component {

	render() {
		const {
			classes
		} = this.props;

		return (
			<ParticlesJS
				className={classes.particles}
				params={particlesConfig}
			/>
		);
	}
}

Particles.propTypes = {
	classes: PropTypes.object.isRequired
};

export default compose(
	withStyles(styles)
)(Particles);
