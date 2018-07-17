import React from "react";
import PropTypes from "prop-types";
import { compose } from "recompose";
// Material UI Imports
import { withStyles } from "@material-ui/core/styles";
import Grow from "@material-ui/core/Grow";
import Fade from "@material-ui/core/Fade";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepLabel from "@material-ui/core/StepLabel";
import StepContent from "@material-ui/core/StepContent";
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CircularProgress from "@material-ui/core/CircularProgress";
// Component Related Imports
import styles from "./styles";
import UsernameForm from "./components/UsernameForm";
import PasswordForm from "./components/PasswordForm";
import PinForm from "./components/PinForm";

function getStepContent(step, next, back, sumbit) {
	switch (step) {
	case 0:
		return <UsernameForm onSubmit={next}/>;
	case 1:
		return (<PasswordForm
			onSubmit={next}
			previousStep={back}
		/>);
	case 2:
		return (<PinForm
			onSubmit={sumbit}
			previousStep={back}
		/>);
	default:
		return "Unknown step";
	}
}

class SetupForm extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			activeStep: 0,
		};
	}

	handleNext = () => {
		this.setState({ activeStep: this.state.activeStep + 1 });
	};

	handleBack = () => {
		this.setState({ activeStep: this.state.activeStep - 1 });
	};

	handleSubmit = (data) => {
		this.handleNext();
		this.props.onSubmit(data);
	}

	render() {
		const { classes } = this.props;
		const { activeStep } = this.state;

		const steps = [
			"Select your username",
			"Create your password",
			"Create your 4-digit PIN"
		];

		return (
			<Grow in>
				<Card className={classes.root}>
					<CardContent>
						<Typography variant="subheading">Create A New Account</Typography>
						<Stepper className={classes.stepper}
							activeStep={activeStep}
							orientation="vertical"
							elevation={0}>
							{steps.map((label, index) => {
								return (
									<Step key={label}>
										<StepLabel>{label}</StepLabel>
										<StepContent>
											{getStepContent(
												index,
												this.handleNext,
												this.handleBack,
												this.handleSubmit
											)}
										</StepContent>
									</Step>
								);
							})}
						</Stepper>
						{activeStep === steps.length && (
							<Fade in>
								<div className={classes.pendingContainer}>
									<CircularProgress className={classes.progress}/>
									<Typography variant="caption"
										className={classes.pendingText}>Your account is being created. Depending on the speed of your device, this process may take a few moments.</Typography>
								</div>
							</Fade>
						)}
					</CardContent>
				</Card>
			</Grow>
		);
	}

}

SetupForm.propTypes = {
	classes: PropTypes.object.isRequired,
	onSubmit: PropTypes.func.isRequired
};

export default compose(
	withStyles(styles)
)(SetupForm);
