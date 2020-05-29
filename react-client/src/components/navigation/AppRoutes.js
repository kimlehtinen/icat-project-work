import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import OtherPage from '../../OtherPage'
import Register from '../authentication/Register'
import Login from '../authentication/Login'
import ProtectedRoute from './ProtectedRoute';

class AppRoutes extends Component {
    render() {
        return (
            <Switch>
                <Route path="/login" render={props => <Login {...props} />} />
                <Route path="/register" render={props => <Register {...props} />} />
                <ProtectedRoute path="/bloodpressure" component={BloodPressure} />
            </Switch>
        );
    }
}


export default AppRoutes;