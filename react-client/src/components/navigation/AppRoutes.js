import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import TemperatureLiveData from '../temperature/TemperatureLiveData'
import TemperatureHistory from '../temperature/TemperatureHistory'
import Dashboard from '../dashboard/Dashboard'
import Register from '../authentication/Register'
import Login from '../authentication/Login'
import ProtectedRoute from './ProtectedRoute'
import { Redirect } from 'react-router-dom'

class AppRoutes extends Component {
    /**
     * This component determines which route/component should be shown to user.
     */

    render() {
        // Go by default to /dashboard
        // If not logged in, user will be redirected to /login
        if (window.location.pathname === '/') {
            return <Redirect to='/dashboard'/>;
        }

        return (
            <Switch>
                <Route path="/login" render={props => <Login {...props} />} />
                <Route path="/register" render={props => <Register {...props} />} />
                <ProtectedRoute path="/dashboard" component={Dashboard} />
                <ProtectedRoute path="/bloodpressure" component={BloodPressure} />
                <ProtectedRoute path="/live-temperature" component={TemperatureLiveData} />
                <ProtectedRoute path="/temperature-history" component={TemperatureHistory} />
            </Switch>
        );
    }
}


export default AppRoutes;