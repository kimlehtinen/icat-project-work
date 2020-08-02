import React, { Component } from 'react';
import { Line } from "react-chartjs-2";
import 'chartjs-plugin-streaming';
import Typography from '@material-ui/core/Typography';
import * as moment from 'moment';


class TemperatureLiveData extends Component {
    _isMounted = false
    ws = new WebSocket('ws://localhost:8080/api/v1/data/current/temperature')
    reference = {};
    state = {
        dataFromServer: null,
        isLoading: true,
        columns: null,
        liveTemperature: '',
        tempData: {
          datasets: [
            {
              label: "Temperature",
              borderColor: "rgb(255, 99, 132)",
              backgroundColor: "rgba(255, 99, 132, 0.5)",
              lineTension: 0,
              borderDash: [8, 4],
              data: []
            }
          ]
        },
        options: {
          scales: {
            yAxes: [{
              scaleLabel: {
                display: true,
                labelString: 'Temperature'
              },
              ticks: {
                beginAtZero: true
              }
            }],
            xAxes: [
              {
                type: "realtime",
                scaleLabel: {
                  display: true,
                  labelString: 'Time'
                }
              }
            ]
          }
        }
    }

    componentDidMount() {
        this._isMounted = true;
 
        this.ws.onopen = () => {
            console.log('connected')
        }

        // recieve data from api websocket
        this.ws.onmessage = evt => {
          if (this._isMounted) {
            this.setState({isLoading: false})

            const data = JSON.parse(evt.data)
            let liveTemperature = 'No live data available (only data < 1min old is shown)';
            
            if (data) {
              const now = moment(new Date());
              const createdAt = moment(data.created_at);
              const timeSinceLast = moment.duration(now.diff(createdAt));
              
              if (timeSinceLast.asMinutes() <= 1) {
                this.state.tempData.datasets[0].data.push({
                  x: Date.now(),
                  y: data.temperature
                });
                liveTemperature = `${data.temperature} °C`;
              }
            }
            
            this.setState({liveTemperature});

            let lineChart = this.reference.chartInstance
            lineChart.update();
          }
        }

        this.ws.onclose = () => {
            console.log('disconnected')
        }

    }

    componentWillUnmount() {
        this._isMounted = false;
        this.ws.close();
    }

    render() {
      return (
        <div>
          <Typography variant="h4" noWrap>
              Live temperature data
          </Typography>
          <br/><br/>
          <Typography variant="h6" noWrap>
          {this.state.liveTemperature}
          </Typography>
          <Line data={this.state.tempData} options={this.state.options} ref = {(reference) => this.reference = reference} />
        </div>
      );
    }
}


export default TemperatureLiveData;