import React from 'react';
import initializeStore from 'modules';
import { Provider } from 'react-redux';
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom';
import { VIndex as VIndexType } from 'models/vindex';
import Root from 'components/Root';
import VideoContent from 'components/VideoContent';

export type AppProps = {
  vindex: VIndexType;
};

export default function App(props: AppProps) {
  const store = initializeStore(props.vindex);

  return (
    <Provider store={store}>
      <BrowserRouter>
        <h1>
          <Link to="/">kinsro</Link>
        </h1>
        <Switch>
          <Route exact path="/">
            <Root />
          </Route>
          <Route path="/contents/:digest">
            <VideoContent />
          </Route>
        </Switch>
      </BrowserRouter>
    </Provider>
  );
}
