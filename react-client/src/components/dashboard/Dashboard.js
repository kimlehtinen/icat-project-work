import React, { Component } from 'react'
import clsx from 'clsx';
import { connect } from 'react-redux';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import { login } from '../../actions/authActions';
import { Redirect } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { Link } from 'react-router-dom';

class Dashboard extends Component {

    static propTypes = {
        isAuthenticated: PropTypes.bool
    }

    componentDidMount() {
        //
    }

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
              <Grid container spacing={3}>
                <Grid item xs={12}>
                  <Typography variant="h3" noWrap>
                      Welcome to MySignals!
                  </Typography>
                  <Typography variant="h5" noWrap>
                      Univeristy of Vaasa
                  </Typography>
                </Grid>
                <Grid item xs={12} sm={4}>
                  <Paper className={classes.paper}>
                    <Link to="/bloodpressure" className={classes.link}>
                      Blood pressure
                    </Link>
                  </Paper>
                </Grid>
                <Grid item xs={12} sm={4}>
                  <Paper className={classes.paper}>
                    <Link to="/live-temperature" className={classes.link}>
                      Live temperature
                    </Link>
                  </Paper>
                </Grid>
                <Grid item xs={12} sm={4}>
                  <Paper className={classes.paper}>
                    <Link to="/temperature-history" className={classes.link}>
                      Temperature history
                    </Link>
                  </Paper>
                </Grid>
              </Grid>
            </div>
        );
    }
}

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
    link: {
      color: 'inherit',
      textDecoration: 'none',
    },
});

const mapStateToProps = state => ({
    isAuthenticated: state.auth.isAuthenticated
});

export default connect(mapStateToProps, null)(withStyles(styles, { withTheme: true })(Dashboard));
