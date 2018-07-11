import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
import classNames from "classnames";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import RadioButtonChecked from "@material-ui/icons/RadioButtonChecked";
import RadioButtonUnchecked from "@material-ui/icons/RadioButtonUnchecked";
// Component Related Imports
import styles from "./styles";

class PinForm extends React.Component {

	constructor(props) {
		super(props);
		this.pinInput = React.createRef();
		this.state = { pin: "" };
	}

	componentDidMount() {
		this.focus();
	}

	focus() {
		this.pinInput.current.focus();
	}

	handleChange(event) {
		this.setState(
			{ pin: event.target.value },
			() => {
				if (this.state.pin.length === 4) {
					this.submit();
				}
			}
		);
	}

	clearPin() {
		this.setState(
			{ pin: "" },
			() => {
				this.pinInput.current.value = "";
			}
		);
	}

	submit() {
		this.props.onSubmit( this.props.username, this.state.pin, this.clearPin.bind(this) );

	}

	render() {
		const {
			classes,
			loading,
			error
		} = this.props;

		let pinDigit = (pinLength, delay) => {
			let className = classNames(
				"kowala",
				[classes.pinDigit],
				{
					animated: loading,
					bounce: loading,
					infinite: loading
				}
			);
			return(
				loading || this.state.pin.length >= pinLength ?
					<RadioButtonChecked
						className={className}
						style={{ animationDelay: delay }}
					/>
					: <RadioButtonUnchecked
						className={className}
					/>
			);
		};

		return (
			<div
				className={classes.root}
				onClick={this.focus.bind(this)}
			>
				<div className={classNames({ "animated": error, "headShake": error })}>
					{pinDigit(1,"0.0s")}
					{pinDigit(2,"0.2s")}
					{pinDigit(3,"0.4s")}
					{pinDigit(4,"0.6s")}
				</div>
				<form>
					<input
						ref={this.pinInput}
						type="number"
						className={classes.input}
						onChange={this.handleChange.bind(this)}
					/>
				</form>
			</div>
		);
	}
}

PinForm.propTypes = {
	classes: PropTypes.object.isRequired,
	onSubmit: PropTypes.func.isRequired,
	loading: PropTypes.bool.isRequired,
	error: PropTypes.bool.isRequired,
	username: PropTypes.string.isRequired
};

export default compose(
	withStyles(styles)
)(PinForm);
