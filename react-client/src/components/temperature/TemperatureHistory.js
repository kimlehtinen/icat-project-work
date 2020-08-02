import React, { Component } from 'react';
import { Line } from "react-chartjs-2";
import 'chartjs-plugin-streaming';
import Typography from '@material-ui/core/Typography';
import * as moment from 'moment';
import { getTemperatureHistoryData } from '../../actions/temperatureActions';
import { withStyles } from '@material-ui/core/styles';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';


class TemperatureHistory extends Component {
    _isMounted = false
    // ws = new WebSocket('ws://localhost:8080/api/v1/data/all/temperature')
    referenceToday = {};
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
          plugins: {
            streaming: false,
          },
          scales: {
            yAxes: [{
              scaleLabel: {
                display: true,
                labelString: 'Temperature'
              }
            }],
            xAxes: [{
              title: "time",
              type: 'time',
              gridLines: {
                  lineWidth: 2
              },
              time: {
                  unit: "minute",
                  unitStepSize: 1000,
                  displayFormats: {
                      millisecond: 'MMM DD',
                      second: 'MMM DD',
                      minute: 'MMM DD',
                      hour: 'MMM DD',
                      day: 'MMM DD',
                      week: 'MMM DD',
                      month: 'MMM DD',
                      quarter: 'MMM DD',
                      year: 'MMM DD',
                  }
              }
          }]
          }
        },
        tempDataToday: {
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
        optionsToday: {
          plugins: {
            streaming: false,
          },
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
            xAxes: [{
              type: 'time',
              time: {
                  unit: 'minute'
              }
            }]
          }
        }
    }

    static propTypes = {
      temperatureHistoryData: PropTypes.array,
      getTemperatureHistoryData: PropTypes.func.isRequired
    }

    componentDidMount() {
        this._isMounted = true;
        this.props.getTemperatureHistoryData();
    }

    componentWillUnmount() {
        this._isMounted = false;
    }

    render() {
      if (this.props.temperatureHistoryData && this.props.temperatureHistoryData.length) {
        this.state.tempDataToday.datasets[0].data = [];
        for (const temp of this.props.temperatureHistoryData) {
          const createdAt = moment(temp.created_at);
          const dataModel = {
            x: createdAt.toDate(),
            y: temp.temperature
          };
          
          // if today
          if (createdAt.isSame(new Date(), "day")) {
            this.state.tempDataToday.datasets[0].data.push(dataModel);
          }
        }
        let lineChart = this.referenceToday.chartInstance
        if (lineChart) {
          lineChart.update();
        }
      }

      return (
        <div>
          <Typography variant="h4" noWrap>
              Temperature data history
          </Typography>
          <br/><br/>
          <Typography variant="h6" noWrap>
          {this.state.liveTemperature}
          </Typography>
          <Line data={this.state.tempDataToday} options={this.state.optionsToday} ref = {(reference) => this.referenceToday = reference} />
        </div>
      );
    }
}

const styles = theme => ({
  //
});

const mapStateToProps = state => ({
  temperatureHistoryData: state.temperature.temperatureHistoryData
});

export default connect(mapStateToProps, {getTemperatureHistoryData})(withStyles(styles, { withTheme: true })(TemperatureHistory));