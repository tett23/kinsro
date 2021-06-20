import React from 'react';
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom';
import { VIndexItem } from 'models/VIndex';
import Root from 'components/Root';
import VideoContent from 'components/VideoContent';
import { RecoilRoot } from 'recoil';
import { mainVindex } from 'modules';

export type AppProps = {
  vindex: VIndexItem[];
};

export default function App(props: AppProps) {
  return (
    <RecoilRoot
      initializeState={({ set }) => {
        set(mainVindex, props.vindex);
      }}
    >
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
    </RecoilRoot>
  );
}
