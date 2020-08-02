import axios from 'axios';
import {
    TEMPERATURE_HISTORY_LOADING,
    TEMPERATURE_HISTORY_LOADED
} from './types';
import { API_VERSION } from '../config';
import { authToken } from '../actions/authActions';

/**
 * Whenever page is loaded, authInit loads user information
 * using jwt token. If this fails, user is redirected to login page.
 * This is fired for example on every refresh in App.js
 */
export const getTemperatureHistoryData = () => (dispatch, getState) => {
  // User loading
  dispatch({ type: TEMPERATURE_HISTORY_LOADING });

  axios
    .get(`/api/v${API_VERSION}/data/all/temperature`, authToken(getState))
    .then(res =>
      dispatch({
        type: TEMPERATURE_HISTORY_LOADED,
        payload: res.data
      })
    )
    .catch(err => {
      // dispatch(returnErrors(err.response.data, err.response.status));
    });
};
