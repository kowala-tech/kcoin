import React from 'react';
import { configure, addDecorator } from '@storybook/react';
import { IntlProvider } from "react-intl";
import MuiThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import muiTheme from "../src/components/MuiTheme";

const withMuiTheme = (story) => (
  <IntlProvider locale="en">
    <MuiThemeProvider theme={muiTheme}>
      { story() }
    </MuiThemeProvider>
  </IntlProvider>
);

// automatically import all files ending in *.stories.js
const req = require.context('../src/components', true, /.stories.js$/);
function loadStories() {
  addDecorator(withMuiTheme)
  req.keys().forEach((filename) => req(filename));
}

configure(loadStories, module);
