import React from 'react';
import initializeStore from 'modules';
import { Provider } from 'react-redux';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { VIndex as VIndexType } from 'models/vindex';
import Root from 'components/Root';

export type AppProps = {
  vindex: VIndexType;
};

export default function App(props: AppProps) {
  const store = initializeStore(props.vindex);

  return (
    <Provider store={store}>
      <BrowserRouter>
        <Switch>
          <Route path="/">
            <Root />
          </Route>
        </Switch>
      </BrowserRouter>
    </Provider>
  );
}
