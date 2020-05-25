import { GET_ERRORS, CLEAR_ERRORS } from './types';

/*
The code in this file is used from
Brad Traversy's "Learn The Mern Stack" Youtube tutorial.
Source code is taken from https://github.com/bradtraversy/mern_shopping_list/tree/master/client (MIT License)
and is modified/built upon by Kim Lehtinen.
*/

// RETURN ERRORS
export const returnErrors = (msg, status, id = null) => {
  return {
    type: GET_ERRORS,
    payload: { msg, status, id }
  };
};

// CLEAR ERRORS
export const clearErrors = () => {
  return {
    type: CLEAR_ERRORS
  };
};