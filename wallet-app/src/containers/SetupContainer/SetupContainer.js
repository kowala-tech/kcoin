import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Material UI Imports
// Component Related Imports
import SetupForm from "../../components/SetupForm";
import { createAccount } from "../../redux/UserRedux/actions";

class SetupContainer extends React.Component {
	render() {
		const { createEdgeAccount } = this.props;

		return (
			<SetupForm onSubmit={createEdgeAccount}/>
		);
	}
}

SetupContainer.propTypes = {
	createEdgeAccount: PropTypes.func.isRequired
};

const mapStateToProps = () => {
	return {
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		createEdgeAccount: (props) => {
			dispatch(createAccount(
				props.username,
				props.password,
				props.pin
			));
		}
	};
};

export default compose(
	connect(mapStateToProps, mapDispatchToProps)
)(SetupContainer);
