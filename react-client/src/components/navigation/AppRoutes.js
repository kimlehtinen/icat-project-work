import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import OtherPage from '../../OtherPage'
import Register from '../authentication/Register'

class AppRoutes extends Component {
    render() {
        return (
            <Switch>
                <Route path="/bloodpressure" component={BloodPressure} />
                <Route path="/" component={Register} />
            </Switch>
        );
    }
}


export default AppRoutes;