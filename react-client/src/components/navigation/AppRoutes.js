import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import OtherPage from '../../OtherPage'
import Register from '../authentication/Register'
import Login from '../authentication/Login'

class AppRoutes extends Component {
    render() {
        return (
            <Switch>
                <Route path="/login" render={props => <Login {...props} />} />
                <Route path="/bloodpressure" component={BloodPressure} />
                <Route path="/register" render={props => <Register {...props} />} />
            </Switch>
        );
    }
}


export default AppRoutes;