import React from "react";
import PropTypes from "prop-types";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
// Component Related Imports
import styles from "./styles";
import CalculatorKey from "./components/CalculatorKey";
import BackSpaceIcon from "@material-ui/icons/Backspace";

const CalculatorOperations = {
	"/": (prevValue, nextValue) => prevValue / nextValue,
	"*": (prevValue, nextValue) => prevValue * nextValue,
	"+": (prevValue, nextValue) => prevValue + nextValue,
	"-": (prevValue, nextValue) => prevValue - nextValue,
	"=": (prevValue, nextValue) => nextValue
};

class CalculatorNumberpad extends React.Component {
	state = {
		value: null,
		displayValue: "0",
		operator: null,
		waitingForOperand: false
	};

	componentDidMount() {
		document.addEventListener("keydown", this.handleKeyDown);
	}

	componentWillUnmount() {
		document.removeEventListener("keydown", this.handleKeyDown);
	}

	clearAll() {
		this.setState({
			value: null,
			displayValue: "0",
			operator: null,
			waitingForOperand: false
		});
		this.props.handleChange("0");
	}

	clearLastChar() {
		const { displayValue } = this.state;
		const newValue = displayValue.substring(0, displayValue.length - 1) || "0";
		this.setState({ displayValue: newValue });
		this.props.handleChange( newValue );
	}

	inputPercent() {
		const { displayValue } = this.state;
		const currentValue = parseFloat(displayValue);

		if (currentValue === 0)
			return;

		const fixedDigits = displayValue.replace(/^-?\d*\.?/, "");
		const parsedValue = parseFloat(displayValue) / 100;
		const newValue = String(parsedValue.toFixed(fixedDigits.length + 2));

		this.setState({ displayValue: newValue });
		this.props.handleChange( newValue );
	}

	inputDot() {
		const { displayValue } = this.state;

		if (!(/\./).test(displayValue)) {
			const newValue = displayValue + ".";
			this.setState({
				displayValue: newValue,
				waitingForOperand: false
			});
			this.props.handleChange( newValue );
		}
	}

	inputDigit(digit) {
		this.props.handleChange(digit);
		const { displayValue, waitingForOperand } = this.state;
		if (waitingForOperand) {
			const newValue = String(digit);
			this.setState({
				displayValue: newValue,
				waitingForOperand: false
			});
			this.props.handleChange( newValue );
		} else {
			const newValue = displayValue === "0" ? String(digit) : displayValue + digit;
			this.setState({ displayValue: newValue });
			this.props.handleChange( newValue );
		}
	}

	performOperation(nextOperator) {
		const { value, displayValue, operator } = this.state;
		const inputValue = parseFloat(displayValue);

		if (value == null) {
			this.setState({
				value: inputValue
			});
		} else if (operator) {
			const currentValue = value || 0;
			const newValue = CalculatorOperations[operator](currentValue, inputValue);
			this.setState({
				value: newValue,
				displayValue: String(newValue)
			});
			this.props.handleChange( newValue );
		}

		this.setState({
			waitingForOperand: true,
			operator: nextOperator
		});
	}

	handleKeyDown = (event) => {
		let { key } = event;

		if (key === "Enter")
			key = "=";

		if ((/\d/).test(key)) {
			event.preventDefault();
			this.inputDigit(parseInt(key, 10));
		} else if (key in CalculatorOperations) {
			event.preventDefault();
			this.performOperation(key);
		} else if (key === ".") {
			event.preventDefault();
			this.inputDot();
		} else if (key === "%") {
			event.preventDefault();
			this.inputPercent();
		} else if (key === "Backspace") {
			event.preventDefault();
			this.clearLastChar();
		} else if (key === "Clear") {
			event.preventDefault();
			this.clearAll();
		}
	};

	render() {
		const { classes } = this.props;

		return (
			<div className={classes.root}>
				<div className={classes.keypad}>
					<div className={classes.inputKeys}>
						<div className={classes.digitKeys}>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDot()}>∙</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(0)}>0</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.clearLastChar()}><BackSpaceIcon /></CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(1)}>1</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(2)}>2</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(3)}>3</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(4)}>4</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(5)}>5</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(6)}>6</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(7)}>7</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(8)}>8</CalculatorKey>
							<CalculatorKey className={classes.digitKey}
								onPress={() => this.inputDigit(9)}>9</CalculatorKey>
						</div>
					</div>
					<div className={classes.operatorKeys}>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("%")}>%</CalculatorKey>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("/")}>÷</CalculatorKey>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("*")}>×</CalculatorKey>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("-")}>−</CalculatorKey>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("+")}>+</CalculatorKey>
						<CalculatorKey className={classes.operatorKey}
							onPress={() => this.performOperation("=")}>=</CalculatorKey>
					</div>
				</div>
			</div>
		);
	}
}

CalculatorNumberpad.propTypes = {
	classes: PropTypes.object.isRequired,
	handleChange: PropTypes.func.isRequired
};

export default withStyles(styles)(CalculatorNumberpad);
