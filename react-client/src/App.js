import React, { Component } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import ResponsiveDrawer from './components/navigation/ResponsiveDrawer'

import { Provider } from 'react-redux';
import store from './store'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

export default function App() {
  const classes = useStyles();

  return (
    <Provider store={store}>
      <ResponsiveDrawer />
    </Provider>
  );
}
