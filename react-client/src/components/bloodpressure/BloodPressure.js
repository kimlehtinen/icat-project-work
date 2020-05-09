import React, { Component } from 'react'
import MaterialTable from 'material-table'

class BloodPressure extends Component {
    _isMounted = false
    ws = new WebSocket('ws://localhost:8080/api/data/all')

    state = {
        dataFromServer: null,
        isLoading: true,
        columns: null,
        data: null
    }

    renderData() {
        if (!(this.state.dataFromServer && this.state.dataFromServer.length)) {
            return (
                <div>No data</div>
            )
        }

        return (
            this.state.dataFromServer.map(function(object, i){
             return <div key={i} >{i}: Diastolic {object.diastolic}, Pulse/Min {object.pulse_per_min}</div>;
            })
        )
    }

    componentDidMount() {
        this.setState({
            columns: [
                // { title: 'Index', field: 'i', type: 'numeric' },
                { title: 'Diastolic', field: 'diastolic', type: 'numeric' },
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
                const message = JSON.parse(evt.data)
                // this.setState({dataFromServer: message})
                this.setState({data: message})
                // console.log(message)
            }
        }

        this.ws.onclose = () => {
            console.log('disconnected')
            // automatically try to reconnect on connection loss
        }

    }

    componentWillUnmount() {
        this._isMounted = false;
    }

    render() {
        if (!this.state.columns || !this.state.data) {
            return (<div>No data</div>)
        }

        return (
            <MaterialTable
            title="Bloodpressure measurements"
            columns={this.state.columns}
            data={this.state.data}
            editable={{
                onRowAdd: (newData) =>
                new Promise((resolve) => {
                    setTimeout(() => {
                    resolve();
                    this.setState((prevState) => {
                        const data = [...prevState.data];
                        data.push(newData);
                        return { ...prevState, data };
                    });
                    }, 600);
                }),
                onRowUpdate: (newData, oldData) =>
                new Promise((resolve) => {
                    setTimeout(() => {
                    resolve();
                    if (oldData) {
                        this.setState((prevState) => {
                        const data = [...prevState.data];
                        data[data.indexOf(oldData)] = newData;
                        return { ...prevState, data };
                        });
                    }
                    }, 600);
                }),
                onRowDelete: (oldData) =>
                new Promise((resolve) => {
                    setTimeout(() => {
                    resolve();
                    this.setState((prevState) => {
                        const data = [...prevState.data];
                        data.splice(data.indexOf(oldData), 1);
                        return { ...prevState, data };
                    });
                    }, 600);
                }),
            }}
            />
        );
    }
}


export default BloodPressure;