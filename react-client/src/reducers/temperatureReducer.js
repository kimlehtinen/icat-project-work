import {
    TEMPERATURE_HISTORY_LOADING,
    TEMPERATURE_HISTORY_LOADED
} from '../actions/types';
  
const initialState = {
    isLoading: false,
    temperatureHistoryData: null
};
  
export default function(state = initialState, action) {
    switch (action.type) {
        case TEMPERATURE_HISTORY_LOADING:
            return {
            ...state,
            isLoading: true
            };
        case TEMPERATURE_HISTORY_LOADED:
            return {
            ...state,
            isLoading: false,
            temperatureHistoryData: action.payload
            };
        default:
            return state;
    }
}
