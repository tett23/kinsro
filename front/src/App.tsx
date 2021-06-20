import React, { useEffect } from 'react';
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom';
import {
  RawVIndexItem,
  sortVIndex,
  toTimestamp,
  VIndexItem,
} from 'models/VIndex';
import Root from 'components/Root';
import VideoContent from 'components/VideoContent';
import { RecoilRoot, useSetRecoilState } from 'recoil';
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

function Effect() {
  const setVIndex = useSetRecoilState(mainVindex);
  useEffect(() => {
    (async () => {
      const vindex = buildVIndex();
      console.time('sort');
      const sorted = await sortVIndex(vindex);
      console.timeEnd('sort');

      setVIndex(sorted);
    })();
  });

  return null;
}

function buildVIndex(): VIndexItem[] {
  console.time('parseJSON');
  const rawVIndex: RawVIndexItem[] = JSON.parse(global.vindex);
  console.timeEnd('parseJSON');

  console.time('transform');
  rawVIndex.forEach((item: RawVIndexItem): VIndexItem => {
    (item as VIndexItem).normalizedFilename = item.filename
      .normalize('NFKC')
      .toLowerCase();
    (item as VIndexItem).timestamp = toTimestamp(item.filename, item.date);

    return item as VIndexItem;
  });
  console.timeEnd('transform');

  return rawVIndex as VIndexItem[];
}
