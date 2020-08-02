import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import TemperatureLiveData from '../temperature/TemperatureLiveData'
import TemperatureHistory from '../temperature/TemperatureHistory'
import OtherPage from '../../OtherPage'
import Register from '../authentication/Register'
import Login from '../authentication/Login'
import ProtectedRoute from './ProtectedRoute';

class AppRoutes extends Component {
    /**
     * This component determines which route/component should be shown to user.
     */

    render() {
        return (
            <Switch>
                <Route path="/login" render={props => <Login {...props} />} />
                <Route path="/register" render={props => <Register {...props} />} />
                <ProtectedRoute path="/bloodpressure" component={BloodPressure} />
                <ProtectedRoute path="/live-temperature" component={TemperatureLiveData} />
                <ProtectedRoute path="/temperature-history" component={TemperatureHistory} />
            </Switch>
        );
    }
}


export default AppRoutes;