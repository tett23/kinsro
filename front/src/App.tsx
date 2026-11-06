import React, { useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { sortVIndex, VIndexItem } from 'models/VIndex';
import Root from 'components/Root';
import VideoContent from 'components/VideoContent';
import { RecoilRoot, useRecoilState } from 'recoil';
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
      <Effect />
      <BrowserRouter>
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

function Effect() {
  const [vindex, setVIndex] = useRecoilState(mainVindex);

  useEffect(() => {
    (async () => {
      console.time('sort');
      const sorted = await sortVIndex(vindex, 'desc');
      console.timeEnd('sort');

      setVIndex(sorted);
    })();
  }, []);

  return null;
}
