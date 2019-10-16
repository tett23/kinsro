import React from 'react';
import initializeStore from 'modules';
import { Provider } from 'react-redux';
import { VIndex as VIndexType } from 'models/vindex';
import VIndex from 'components/VIndex';

export type AppProps = {
  vindex: VIndexType;
};

export default function App(props: AppProps) {
  const store = initializeStore(props.vindex);
  console.log(VIndex);

  return (
    <Provider store={store}>
      <VIndex />
    </Provider>
  );
}
