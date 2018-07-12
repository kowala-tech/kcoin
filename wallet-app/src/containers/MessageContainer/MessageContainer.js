import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Component Related Imports
import Snackbar from "../../components/Snackbar";
import { closeMessage } from "../MessageContainer/redux/actions";

class MessageContainer extends React.Component {
	render() {
		const {
			text,
			open,
			handleClose
		} = this.props;

		return (<Snackbar
			open={open}
			message={text}
			handleClose={handleClose}
		/>);
	}
}

MessageContainer.propTypes = {
	text: PropTypes.string,
	open: PropTypes.bool.isRequired,
	handleClose: PropTypes.func
};

const mapStateToProps = (state) => {
	return {
		text: state.ui.message.text,
		open: state.ui.message.open
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		handleClose: () => {
			dispatch(closeMessage());
		}
	};
};

export default compose(
	connect(mapStateToProps, mapDispatchToProps),
)(MessageContainer);
