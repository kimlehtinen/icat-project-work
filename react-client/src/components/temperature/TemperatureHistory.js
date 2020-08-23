import React, { Component } from 'react';
import { Line } from "react-chartjs-2";
import 'chartjs-plugin-streaming';
import Typography from '@material-ui/core/Typography';
import * as moment from 'moment';
import { getTemperatureHistoryData } from '../../actions/temperatureActions';
import { withStyles } from '@material-ui/core/styles';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';
import ArrowForwardIosIcon from '@material-ui/icons/ArrowForwardIos';

class TemperatureHistory extends Component {
    _isMounted = false
    
    referenceToday = {};
    referenceThisWeek = {};
    referenceMonth = {};
    
    TAB_DAY = 0;
    TAB_WEEK = 1;
    TAB_MONTH = 2;

    state = {
      currentDatesSet: false,
      tabIndex: 0,
      dayChart: {
        currentDate: null,
      },
      weekChart: {
        currentWeekDate: null,
        currentWeek: null,
        currentYear: null,
      },
      monthChart: {
        currentDate: null,
        currentMonth: null,
        currentYear: null,
      },
      // temperature today line data
      tempDataToday: {
          datasets: [
            {
              label: "Temperature today",
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
                labelString: 'Temperature today'
              },
              ticks: {
                beginAtZero: true
              }
            }],
            xAxes: [{
              type: 'time',
              time: {
                  unit: 'hour'
              }
            }]
          }
      },
      // temperature this week line data
      tempDataThisWeek: {
          datasets: [
            {
              label: "Temperature this week",
              borderColor: "rgb(255, 99, 132)",
              backgroundColor: "rgba(255, 99, 132, 0.5)",
              lineTension: 0,
              borderDash: [8, 4],
              data: []
            }
          ]
      },
      optionsThisWeek: {
          plugins: {
            streaming: false,
          },
          scales: {
            yAxes: [{
              scaleLabel: {
                display: true,
                labelString: 'Temperature this week'
              },
              ticks: {
                beginAtZero: true
              }
            }],
            xAxes: [{
              title: "time",
              type: 'time',
              gridLines: {
                  lineWidth: 2
              },
              time: {
                  unit: "day",
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
      // temperature this month line data
      tempDataMonth: {
          datasets: [
            {
              label: "Temperature this month",
              borderColor: "rgb(255, 99, 132)",
              backgroundColor: "rgba(255, 99, 132, 0.5)",
              lineTension: 0,
              borderDash: [8, 4],
              data: []
            }
          ]
      },
      optionsMonth: {
          plugins: {
            streaming: false,
          },
          scales: {
            yAxes: [{
              scaleLabel: {
                display: true,
                labelString: 'Temperature this month'
              },
              ticks: {
                beginAtZero: true
              }
            }],
            xAxes: [{
              title: "time",
              type: 'time',
              gridLines: {
                  lineWidth: 2
              },
              time: {
                  unit: "month",
                  unitStepSize: 1000,
                  displayFormats: {
                      millisecond: 'MMM',
                      second: 'MMM',
                      minute: 'MMM',
                      hour: 'MMM',
                      day: 'MMM',
                      week: 'MMM',
                      month: 'MMM',
                      quarter: 'MMM',
                      year: 'MMM',
                  }
              }
          }]
        }
      }
    }

    static propTypes = {
      temperatureHistoryData: PropTypes.array,
      getTemperatureHistoryData: PropTypes.func.isRequired,
    }

    componentDidMount() {
        this._isMounted = true;
        this.props.getTemperatureHistoryData();
        const now = moment(new Date());
        const currentWeek = now.isoWeek();
        const currentDate = now;
        const currentYear = now.isoWeekYear();
        const currentMonth = now.month();
        
        this.setState({ 
          dayChart: { currentDate },
          weekChart: { currentWeek, currentWeekDate: currentDate, currentYear},
          monthChart: { currentDate, currentMonth, currentYear }
        });
    }

    componentWillUnmount() {
        this._isMounted = false;
    }

    populateTemperatureDay(day) {
      this.state.tempDataToday.datasets[0].data = [];
      for (const temp of this.props.temperatureHistoryData) {
        const createdAt = moment(temp.created_at);
        const dataModel = {
          x: createdAt.toDate(),
          y: temp.temperature
        };
          
        if (createdAt.isSame(moment(day), 'day')) {
          this.state.tempDataToday.datasets[0].data.push(dataModel);
        }
      }
      const lineChart = (this.referenceToday && this.referenceToday.chartInstance) ? this.referenceToday.chartInstance : null;
      if (lineChart) {
        lineChart.update();
      }
    }

    populateTemperatureWeek(year, week) {
      if (!(this.state.weekChart && this.state.weekChart.currentYear && this.state.weekChart.currentWeek)) {
        return;
      }

      this.state.tempDataThisWeek.datasets[0].data = [];
      for (const temp of this.props.temperatureHistoryData) {
        const createdAt = moment(temp.created_at);
        const dataModel = {
          x: createdAt.format('MMM DD'),
          y: temp.temperature
        };
          
        if (createdAt.isoWeekYear() == year && createdAt.isoWeek() == week) {
          this.state.tempDataThisWeek.datasets[0].data.push(dataModel);
        }
      }
      const lineChart = (this.referenceThisWeek && this.referenceThisWeek.chartInstance) ? this.referenceThisWeek.chartInstance : null;
      if (lineChart) {
        console.log('Update chart');
        lineChart.update();
      }
    }

    populateTemperatureMonth(year, month) {
      if (!(this.state.monthChart && this.state.monthChart.currentYear && this.state.monthChart.currentMonth)) {
        return;
      }

      this.state.tempDataMonth.datasets[0].data = [];
      for (const temp of this.props.temperatureHistoryData) {
        console.log(temp);
        const createdAt = moment(temp.created_at);
        const dataModel = {
          x: createdAt.format('MMM DD'),
          y: temp.temperature
        };
        if (createdAt.isoWeekYear() == year && createdAt.month() == month) {
          this.state.tempDataMonth.datasets[0].data.push(dataModel);
        }
      }
      const lineChart = (this.referenceMonth && this.referenceMonth.chartInstance) ? this.referenceMonth.chartInstance : null;
      if (lineChart) {
        lineChart.update();
      }
    }

    prevDay = () => {
      const prevDay = moment(this.state.dayChart.currentDate.toDate());
      prevDay.subtract(1, 'day');
      this.setState({
        dayChart: {
          currentDate: prevDay
        }
      });
      this.populateTemperatureDay(prevDay.toDate());
    };

    nextDay = () => {
      const nextDay = moment(this.state.dayChart.currentDate.toDate());
      nextDay.add(1, 'day');
      const now = moment(new Date());

      if (!nextDay.isAfter(now, 'day')) {
        this.setState({
          dayChart: {
            currentDate: nextDay
          }
        });
        this.populateTemperatureDay(nextDay.toDate());
      }
    };
    
    prevWeek = () => {
      const prevWeekDate = moment(this.state.weekChart.currentWeekDate.toDate());
      prevWeekDate.subtract(1, 'week');
      this.setState({
        weekChart: {
          currentWeek: prevWeekDate.isoWeek(), 
          currentWeekDate: prevWeekDate, 
          currentYear: prevWeekDate.isoWeekYear()
        }
      });
      this.populateTemperatureWeek(prevWeekDate.isoWeekYear(), prevWeekDate.isoWeek());
    };

    nextWeek = () => {
      const nextWeekDate = moment(this.state.weekChart.currentWeekDate.toDate());
      nextWeekDate.add(1, 'week');
      const now = moment(new Date());
      if (!(nextWeekDate.isoWeekYear() >= now.isoWeekYear() && nextWeekDate.isoWeek() > now.isoWeek())) {
        this.setState({
          weekChart: {
            currentWeek: nextWeekDate.isoWeek(), 
            currentWeekDate: nextWeekDate, 
            currentYear: nextWeekDate.isoWeekYear()
          }
        });
        this.populateTemperatureWeek(nextWeekDate.isoWeekYear(), nextWeekDate.isoWeek());
      }
    };

    prevMonth = () => {
      const prevMonthDate = moment(this.state.monthChart.currentDate.toDate());
      prevMonthDate.subtract(1, 'month');
      this.setState({
        monthChart: {
          currentMonth: prevMonthDate.month(), 
          currentDate: prevMonthDate, 
          currentYear: prevMonthDate.isoWeekYear()
        }
      });
      this.populateTemperatureMonth(prevMonthDate.isoWeekYear(), prevMonthDate.month());
    };

    nextMonth = () => {
      const nextMonthDate = moment(this.state.monthChart.currentDate.toDate());
      nextMonthDate.add(1, 'month');
      const now = moment(new Date());
      if (!(nextMonthDate.isoWeekYear() >= now.isoWeekYear() && nextMonthDate.month() > now.month())) {
        this.setState({
          monthChart: {
            currentMonth: nextMonthDate.month(), 
            currentDate: nextMonthDate, 
            currentYear: nextMonthDate.isoWeekYear()
          }
        });
        this.populateTemperatureMonth(nextMonthDate.isoWeekYear(), nextMonthDate.month());
      }
    };

    tabProps(index) {
      return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
      };
    }

    handleTabClick = (event, value) => {
      this.setState({tabIndex: value});
    };

    setCurrentDates() {
      if (!this.state.currentDatesSet && this.props.temperatureHistoryData && this.props.temperatureHistoryData.length) {
        const now = moment(new Date());
        const currentWeek = now.isoWeek();
        const currentYear = now.isoWeekYear();
        const currentMonth = now.month();

        this.populateTemperatureDay(now.toDate());
        this.populateTemperatureWeek(currentYear, currentWeek);
        this.populateTemperatureMonth(currentYear, currentMonth);
        this.setState({currentDatesSet: true});
      }
    }

    render() {
      const { classes } = this.props;
      this.setCurrentDates();

      return (
        <div>
          <Typography variant="h4" noWrap>
              Temperature data history
          </Typography>
          <br/><br/>
          <div className={classes.root}>
            <AppBar position="static">
              <Tabs value={this.state.tabIndex} onChange={this.handleTabClick} aria-label="tabs" centered>
                <Tab label="DAY" {...this.tabProps(this.TAB_DAY)}/>
                <Tab label="WEEK" {...this.tabProps(this.TAB_WEEK)}/>
                <Tab label="MONTH" {...this.tabProps(this.TAB_MONTH)} />
              </Tabs>
            </AppBar>
            <div
              role="tabpanel"
              hidden={this.state.tabIndex !== this.TAB_DAY}
              id={`simple-tabpanel-${this.TAB_DAY}`}
              aria-labelledby={`simple-tab-${this.TAB_DAY}`}
            >
              {this.state.tabIndex === this.TAB_DAY && (
                <Box p={3}>
                  <Grid container direction="row" justify="center" alignItems="center" spacing={3}> 
                    <Grid item xs={12} className={classes.textCenter}>
                      {this.state.dayChart.currentDate
                      && 
                      <Typography variant="h4" noWrap>
                      <Button variant="contained" onClick={this.prevDay}>
                        <ArrowBackIosIcon />
                      </Button>
                      Day: {moment(this.state.dayChart.currentDate).format('MMM DD YYYY')}
                      <Button variant="contained" onClick={this.nextDay}>
                        <ArrowForwardIosIcon />
                      </Button>
                      </Typography>}
                    </Grid> 
                  </Grid>
                  <Line data={this.state.tempDataToday} options={this.state.optionsToday} ref = {(reference) => this.referenceToday = reference} />
                </Box>
              )}
            </div>
            <div
              role="tabpanel"
              hidden={this.state.tabIndex !== this.TAB_WEEK}
              id={`simple-tabpanel-${this.TAB_WEEK}`}
              aria-labelledby={`simple-tab-${this.TAB_WEEK}`}
            >
              {this.state.tabIndex === this.TAB_WEEK && (
                <Box p={3}>
                  <Grid container direction="row" justify="center" alignItems="flex-start" spacing={3}>
                    <Grid item xs={12} className={classes.textCenter}>
                      {(this.state.weekChart.currentWeek && this.state.weekChart.currentYear) 
                      && 
                      <Typography variant="h4" noWrap>
                      <Button variant="contained" onClick={this.prevWeek}>
                        <ArrowBackIosIcon />
                      </Button>
                      Week {this.state.weekChart.currentWeek}
                      ({this.state.weekChart.currentYear})
                      <Button variant="contained" onClick={this.nextWeek}>
                        <ArrowForwardIosIcon />
                      </Button>
                      </Typography>}
                    </Grid>
                  </Grid>
                  <Line data={this.state.tempDataThisWeek} options={this.state.optionsThisWeek} ref = {(reference) => this.referenceThisWeek = reference} />
                </Box>
              )}
            </div>
            <div
              role="tabpanel"
              hidden={this.state.tabIndex !== this.TAB_MONTH}
              id={`simple-tabpanel-${this.TAB_MONTH}`}
              aria-labelledby={`simple-tab-${this.TAB_MONTH}`}
            >
              {this.state.tabIndex === this.TAB_MONTH && (
              <Box p={3}>
                <Grid container direction="row" justify="center" alignItems="flex-start" spacing={3}>
                  <Grid item xs={12} className={classes.textCenter}>
                    {(this.state.monthChart.currentMonth && this.state.monthChart.currentYear) 
                    && 
                    <Typography variant="h4" noWrap>
                    <Button variant="contained" onClick={this.prevMonth}>
                      <ArrowBackIosIcon />
                    </Button>
                    Month: {moment(this.state.monthChart.currentDate).format('MMM')}
                    {` (${this.state.monthChart.currentYear})`}
                    <Button variant="contained" onClick={this.nextMonth}>
                      <ArrowForwardIosIcon />
                    </Button>
                    </Typography>}
                  </Grid>
                </Grid>
                <Line data={this.state.tempDataMonth} options={this.state.optionsMonth} ref = {(reference) => this.referenceMonth = reference} />
              </Box>
              )}
            </div>
          </div>
        </div>
      );
    }
}

const styles = theme => ({
  root: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
  },
  textCenter: {
    textAlign: 'center',
  },
});

const mapStateToProps = state => ({
  temperatureHistoryData: state.temperature.temperatureHistoryData
});

export default connect(mapStateToProps, {getTemperatureHistoryData})(withStyles(styles, { withTheme: true })(TemperatureHistory));
