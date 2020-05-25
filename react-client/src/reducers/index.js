import { combineReducers } from 'redux'
import authReducer from './authReducer';
import errorReducer from './errorReducer';

/*
The code in this file is used from
Brad Traversy's "Learn The Mern Stack" Youtube tutorial.
Source code is taken from https://github.com/bradtraversy/mern_shopping_list/tree/master/client (MIT License)
and is modified/built upon by Kim Lehtinen.
*/

export default combineReducers({
    error: errorReducer,
    auth: authReducer
});