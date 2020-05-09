import React, { Component } from 'react';
import {
    Switch,
    Route
  } from 'react-router-dom';
import BloodPressure from '../bloodpressure/BloodPressure'
import OtherPage from '../../OtherPage'

class AppRoutes extends Component {
    render() {
        return (
            <Switch>
                <Route path="/bloodpressure" component={BloodPressure} />
                <Route path="/" component={OtherPage} />
            </Switch>
        );
    }
}


export default AppRoutes;