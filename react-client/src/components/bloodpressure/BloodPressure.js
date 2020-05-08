import React, { Component } from 'react';

class BloodPressure extends Component {
    ws = new WebSocket('ws://localhost:8080/api/data/all');

    state = {
        dataFromServer: null
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
        this.ws.onopen = () => {
            // on connecting, do nothing but log it to the console
            console.log('connected')
        }

        this.ws.onmessage = evt => {
            // listen to data sent from the websocket server
            const message = JSON.parse(evt.data)
            this.setState({dataFromServer: message})
            console.log(message)
        }

        this.ws.onclose = () => {
            console.log('disconnected')
            // automatically try to reconnect on connection loss
        }

    }

    render() {
        return (
            <div>
                {this.renderData()}
            </div>
        );
    }
}


export default BloodPressure;