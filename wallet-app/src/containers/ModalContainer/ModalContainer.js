import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Component Related Imports
import SendModal from "../SendModal";
import DepositModal from "../DepositModal";

const MODAL_COMPONENTS = {
	"SEND_MODAL": SendModal,
	"DEPOSIT_MODAL": DepositModal
};

class ModalContainer extends React.Component {
	render() {
		if (!this.props.modalType) { return null; }
		const ModalToRender = MODAL_COMPONENTS[this.props.modalType];
		return <ModalToRender />;
	}
}

ModalContainer.propTypes = {
	modalType: PropTypes.string
};

const mapStateToProps = (state) => {
	return {
		modalType: state.ui.modal.type
	};
};

export default compose(
	connect(mapStateToProps),
)(ModalContainer);
