import React, { useEffect } from 'react';
import ResponsiveDrawer from './components/navigation/ResponsiveDrawer'
import { Provider } from 'react-redux';
import store from './store'
import { authInit } from './actions/authActions';


const App = () => {
  useEffect(() => {
    store.dispatch(authInit());
  }, []);

  return (
    <Provider store={store}>
      <ResponsiveDrawer />
    </Provider>
  );
}

export default App;
