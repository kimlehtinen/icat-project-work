import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { logout } from '../../actions/authActions';
import PropTypes from 'prop-types';

class Logout extends Component {
    static propTypes = {
        logout: PropTypes.func.isRequired
    };

    render() {
        return (
            <Fragment>
                <div onClick={this.props.logout} href="#">Logout</div>
            </Fragment>
        );
    }
}

export default connect(null, { logout })(Logout);