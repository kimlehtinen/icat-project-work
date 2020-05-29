import axios from 'axios';
import { returnErrors } from './errorActions';
import {
  USER_LOADED,
  USER_LOADING,
  AUTH_ERROR,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT_SUCCESS,
  REGISTER_SUCCESS,
  REGISTER_FAIL
} from './types';

/*
The code in this file is used from
Brad Traversy's "Learn The Mern Stack" Youtube tutorial.
Source code is taken from https://github.com/bradtraversy/mern_shopping_list/tree/master/client (MIT License)
and is modified/built upon by Kim Lehtinen.
*/

// Check token & load user
export const authInit = () => (dispatch, getState) => {
  // User loading
  dispatch({ type: USER_LOADING });

  axios
    .get('/api/v1/auth/user', tokenConfig(getState))
    .then(res =>
      dispatch({
        type: USER_LOADED,
        payload: res.data
      })
    )
    .catch(err => {
      // dispatch(returnErrors(err.response.data, err.response.status));
      dispatch({
        type: AUTH_ERROR
      });
    });
};

// Register User
export const register = ({ email, password }) => (
  dispatch
) => {
  console.log('Inside register action');

  // Headers
  const config = {
    headers: {
      'Content-Type': 'application/json'
    }
  };
  // Request body
  const body = JSON.stringify({ email, password });

  axios
    .post('/api/v1/auth/register', body, config)
    .then(res => {
      console.log('RES:', res)
      dispatch({
        type: REGISTER_SUCCESS,
        payload: res.data
      })
    }
    )
    .catch(err => {
      console.log('ERROR:', err);
      /*dispatch(
        returnErrors(err.response.data, err.response.status, 'REGISTER_FAIL')
      );
      dispatch({
        type: REGISTER_FAIL
      });*/
    });
};

// Login User
export const login = ({ email, password }) => (
  dispatch
) => {
  // Headers
  const config = {
    headers: {
      'Content-Type': 'application/json'
    },
    crossDomain: true
  };

  // Request body
  const body = JSON.stringify({ email, password });

  axios
    .post('/api/v1/auth/login', body, config)
    .then(res =>
      dispatch({
        type: LOGIN_SUCCESS,
        payload: res.data
      })
    )
    .catch(err => {
      /*dispatch(
        returnErrors(err.response.data, err.response.status, 'LOGIN_FAIL')
      );*/
      console.log('LOGIN ERROR:', err);
      dispatch({
        type: LOGIN_FAIL
      });
    });
};

// Logout User
export const logout = () => {
  return {
    type: LOGOUT_SUCCESS
  };
};

// Setup config/headers and token
export const tokenConfig = (getState) => {
  // Get token from localstorage
  const token = getState().auth.token;

  // Headers
  const config = {
    headers: {
      'Content-type': 'application/json'
    }
  };

  if (token) {
    // add bearer to http headers
    config.headers['Authorization'] = `Bearer ${token}`;
  }

  return config;
};