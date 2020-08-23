import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import Logout from '../authentication/Logout';
import { Link } from 'react-router-dom';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import HomeIcon from '@material-ui/icons/Home';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import PersonAddIcon from '@material-ui/icons/PersonAdd';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import DataUsageIcon from '@material-ui/icons/DataUsage';

class SideMenu extends Component {

    static propTypes = {
        auth: PropTypes.object.isRequired
    }

    isAuthenticated() {
        const  { auth } = this.props;
        return auth && auth.isAuthenticated;
    }

    menuItems() {
        const { classes } = this.props; 
        
        const loggedInMenuItems = (
            <Fragment>
                <Link to="/dashboard" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><HomeIcon /></ListItemIcon>
                        <ListItemText primary={'Dashboard'} />
                    </ListItem>
                </Link>
                <Link to="/bloodpressure" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><DataUsageIcon /></ListItemIcon>
                        <ListItemText primary={'Bloodpressure data'} />
                    </ListItem>      
                </Link>
                <Link to="/live-temperature" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><DataUsageIcon /></ListItemIcon>
                        <ListItemText primary={'Live temperature data'} />
                    </ListItem>      
                </Link>
                <Link to="/temperature-history" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><DataUsageIcon /></ListItemIcon>
                        <ListItemText primary={'Temperature history'} />
                    </ListItem>      
                </Link>
                <ListItem button>
                    <ListItemIcon><ExitToAppIcon /></ListItemIcon>
                    <ListItemText><Logout /></ListItemText>
                </ListItem>
            </Fragment>
        );

        const notLoggedInMenuItems = (
            <Fragment>
                <Link to="/login" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><AccountCircleIcon /></ListItemIcon>
                        <ListItemText primary={'Login'} />
                    </ListItem>
                </Link>
                <Link to="/register" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><PersonAddIcon /></ListItemIcon>
                        <ListItemText primary={'Register'} />
                    </ListItem>
                </Link>
            </Fragment>
        )

        if (this.isAuthenticated()) {
            return loggedInMenuItems;
        }

        return notLoggedInMenuItems;
    }

    render() {
        return (
            <List>
                { this.menuItems() }
            </List>
        );
    }
}

const styles = theme => ({
    link: {
        color: 'inherit',
        textDecoration: 'none',
    },
});

const mapStateToProps = state => ({
    auth: state.auth
});

export default connect(mapStateToProps, null)(withStyles(styles, { withTheme: true })(SideMenu));
