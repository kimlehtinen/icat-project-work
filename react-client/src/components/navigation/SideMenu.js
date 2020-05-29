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
import InboxIcon from '@material-ui/icons/MoveToInbox';

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
                <Link to="/bloodpressure" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><InboxIcon /></ListItemIcon>
                        <ListItemText primary={'Bloodpressure data'} />
                    </ListItem>
                    
                </Link>
                <ListItem button>
                    <ListItemIcon><InboxIcon /></ListItemIcon>
                    <ListItemText><Logout /></ListItemText>
                </ListItem>
            </Fragment>
        );

        const notLoggedInMenuItems = (
            <Fragment>
                <ListItem button>
                    <ListItemIcon><InboxIcon /></ListItemIcon>
                    <ListItemText primary={'Guest Link'} />
                </ListItem>
                <Link to="/register" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><InboxIcon /></ListItemIcon>
                        <ListItemText primary={'Register'} />
                    </ListItem>
                </Link>
                <Link to="/login" className={classes.link}>
                    <ListItem button>
                        <ListItemIcon><InboxIcon /></ListItemIcon>
                        <ListItemText primary={'Login'} />
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
