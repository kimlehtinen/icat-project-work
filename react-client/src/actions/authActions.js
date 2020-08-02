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

import { API_VERSION } from '../config';

/*
The code in this file is used from
Brad Traversy's "Learn The Mern Stack" Youtube tutorial.
Source code is taken from https://github.com/bradtraversy/mern_shopping_list/tree/master/client (MIT License)
and is modified/built upon by Kim Lehtinen.
*/


/**
 * Whenever page is loaded, authInit loads user information
 * using jwt token. If this fails, user is redirected to login page.
 * This is fired for example on every refresh in App.js
 */
export const authInit = () => (dispatch, getState) => {
  // User loading
  dispatch({ type: USER_LOADING });

  axios
    .get(`/api/v${API_VERSION}/auth/user`, authToken(getState))
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

/**
 * Registers new user
 * 
 * @param {*} newUser{email, password}
 */
export const register = ({ email, password }) => (
  dispatch
) => {

  // http headers
  const config = {
    headers: {
      'Content-Type': 'application/json'
    }
  };

  // POST request body
  const body = JSON.stringify({ email, password });

  axios
    .post(`/api/v${API_VERSION}/auth/register`, body, config)
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

/**
 * Logs in user
 * 
 * @param {*} user{email, password}
 */
export const login = ({ email, password }) => (
  dispatch
) => {
  // http headers
  const config = {
    headers: {
      'Content-Type': 'application/json'
    },
    crossDomain: true
  };

  // Request body
  const body = JSON.stringify({ email, password });

  axios
    .post(`/api/v${API_VERSION}/auth/login`, body, config)
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

/**
 * Logs out user
 */
export const logout = () => {
  return {
    type: LOGOUT_SUCCESS
  };
};

/**
 * Add jwt token to http headers
 * 
 * @param {*} getState 
 */
export const authToken = (getState) => {
  // Get token from localstorage
  const token = getState().auth.token;

  // http headers
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