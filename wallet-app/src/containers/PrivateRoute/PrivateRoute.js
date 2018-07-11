import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { compose } from "recompose";
// Component Related Imports1
import { Route, Redirect } from "react-router-dom";

class PrivateRoute extends React.Component {
	render() {
		const {
			authenticated,
			component: Component,
			...props
		} = this.props;

		return (
			<Route
				{...props}
				render={props =>
					authenticated
						? <Component {...props} />
						: <Redirect to={{ pathname: "/login", state: { from: props.location } }} />
				}
			/>
		);
	}
}

PrivateRoute.propTypes = {
	authenticated: PropTypes.bool.isRequired,
	component: PropTypes.func.isRequired
};

const mapStateToProps = (state) => {
	return {
		authenticated: state.user.authenticated
	};
};

export default compose(
	connect(mapStateToProps, {})
)(PrivateRoute);
