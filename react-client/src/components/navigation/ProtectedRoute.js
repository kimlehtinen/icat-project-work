import React, { Component } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

class ProtectedRoute extends Component {
    /**
     * This component checks if a user is allowed to visit a route or not.
     * If not authenticated, user is redirected to login page.
     */

    static propTypes = {
        isAuthenticated: PropTypes.bool
    }

    render() {
       const {component: Component, ...rest} = this.props;

       if (this.props.isAuthenticated === false) {
        return (
            <Route {...rest} render={props => (
                <Redirect to={{
                    pathname: '/login', 
                    state: {from: props.location }
                }}/>
             )}/>
        )
       }

        return (
            <Route {...rest} render={props => (
                <Component {...props}/>
            )}/>
        );
    }
}

const mapStateToProps = state => ({
    isAuthenticated: state.auth.isAuthenticated,
});

export default connect(mapStateToProps, null)(ProtectedRoute);
