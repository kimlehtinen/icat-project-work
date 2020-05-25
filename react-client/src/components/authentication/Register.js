import React, { Component } from 'react'
import clsx from 'clsx';
import IconButton from '@material-ui/core/IconButton';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import FormControl from '@material-ui/core/FormControl';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import EmailIcon from '@material-ui/icons/Email';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { connect } from 'react-redux';
import { withStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';
import { register } from '../../actions/authActions';

class Register extends Component {

    state = {
        email: '',
        password: '',
        showPassword: false,
        msg: null,
    }

    static propTypes = {
        isAuthenticated: PropTypes.bool,
        error: PropTypes.object.isRequired,
        register: PropTypes.func.isRequired
    }

    componentDidMount() {
        //
    }

    handleChange = (prop) => (event) => {
        this.setState({ [prop]: event.target.value })
    };

    handleClickShowPassword = () => {
        this.setState({ showPassword: !this.state.showPassword })
    };

    handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    handleClickRegister = () => {
        console.log('Registering...');
        const newUser = {
            email: this.state.email,
            password: this.state.password
        };

        // attempt register
        this.props.register(newUser)
    };

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
                <Paper className={classes.paper}>
                    <Grid container spacing={3}>
                        <Grid item xs={12}>
                            <FormControl fullWidth className={clsx(classes.margin, classes.textField)}>
                            <InputLabel htmlFor="standard-adornment-email">Email</InputLabel>
                            <Input
                                id="standard-adornment-email"
                                type={'text'}
                                value={this.state.email}
                                onChange={this.handleChange('email')}
                                endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                    aria-label="email address"
                                    >
                                    {<EmailIcon />}
                                    </IconButton>
                                </InputAdornment>
                                }
                            />
                            </FormControl>
                        </Grid>
                        <Grid item xs={12}>
                            <FormControl fullWidth className={clsx(classes.margin, classes.textField)}>
                            <InputLabel htmlFor="standard-adornment-password">Password</InputLabel>
                            <Input
                                id="standard-adornment-password"
                                type={this.state.showPassword ? 'text' : 'password'}
                                value={this.state.password}
                                onChange={this.handleChange('password')}
                                endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={this.handleClickShowPassword}
                                    onMouseDown={this.handleMouseDownPassword}
                                    >
                                    {this.state.showPassword ? <Visibility /> : <VisibilityOff />}
                                    </IconButton>
                                </InputAdornment>
                                }
                            />
                            </FormControl>
                        </Grid>
                        <Grid item xs={12}>
                        <FormControl fullWidth className={clsx(classes.margin)}>
                            <Button 
                            variant="contained" 
                            color="primary"
                            onClick={this.handleClickRegister}
                            >
                            Register
                            </Button>
                        </FormControl>
                        </Grid>
                    </Grid>
                </Paper>
            </div>
        );
    }
}

const styles = theme => ({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
});

const mapStateToProps = state => ({
    isAuthenticated: state.auth.isAuthenticated,
    error: state.error
});

export default connect(mapStateToProps, {register})(withStyles(styles, { withTheme: true })(Register));
