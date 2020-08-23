import React, { Component } from 'react'
import MaterialTable from 'material-table'
import CircularProgress from '@material-ui/core/CircularProgress'
import { withStyles } from '@material-ui/core/styles'
import { connect } from 'react-redux'

class BloodPressure extends Component {
    _isMounted = false
    ws = new WebSocket('ws://localhost:8080/api/v1/data/all-live/blood-pressure')

    state = {
        dataFromServer: null,
        isLoading: true,
        columns: null,
        data: null
    }

    componentDidMount() {
        this.setState({
            columns: [
                // { title: 'Index', field: 'i', type: 'numeric' },
                { title: 'Date', field: 'created_at' },
                { title: 'Diastolic', field: 'diastolic', type: 'numeric' },
                { title: 'Systolic', field: 'systolic', type: 'numeric' },
                { title: 'Pulse/Min', field: 'pulse_per_min', type: 'numeric' },
            ]
        })

        this._isMounted = true;
 
        this.ws.onopen = () => {
            // on connecting, do nothing but log it to the console
            console.log('connected')
        }

        this.ws.onmessage = evt => {
            if (this._isMounted) {
                this.setState({isLoading: false})
                // listen to data sent from the websocket server
                const bpData = JSON.parse(evt.data)
                // this.setState({dataFromServer: message})
                this.setState({data: bpData})
            }
        }

        this.ws.onclose = () => {
            console.log('disconnected')
        }

    }

    componentWillUnmount() {
        this._isMounted = false;
    }

    render() {
        const { classes } = this.props;

        if (this.state.isLoading) {
            return (
                <div className={classes.textCenter}>
                    <CircularProgress />
                </div>
            );
        }

        if (!this.state.columns || !this.state.data) {
            return (<div>No data</div>)
        }

        return (
            <MaterialTable
            title="Bloodpressure measurements"
            columns={this.state.columns}
            data={this.state.data}
            />
        );
    }
}

const styles = theme => ({
    textCenter: {
      textAlign: 'center',
    },
});

const mapStateToProps = state => ({
    //
});

export default connect(mapStateToProps, null)(withStyles(styles, { withTheme: true })(BloodPressure));