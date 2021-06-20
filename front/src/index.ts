import 'core-js/stable';
import 'regenerator-runtime/runtime';
import './wasm_exec';
import React from 'react';
import { render } from 'react-dom';
import App from './App';
import { RawVIndexItem, VIndexItem } from 'models/vindex';
import { VIndexItem } from 'components/VIndexItem';

(async () => {
  const go = new Go();
  const buf = await fetch('/kinsro.wasm').then((resp) => resp.arrayBuffer());
  const result = await WebAssembly.instantiate(buf, go.importObject).catch(
    (err: Error) => err,
  );
  if (result instanceof Error) {
    console.error(result);
    return;
  }

  go.run(result.instance);
  const vindexResult = await fetch('/vindex');
  const bytes = new Uint8Array(await vindexResult.arrayBuffer());

  setTimeout(async () => {
    window.parseVIndex(bytes);
    const vindex = JSON.parse(window.vindex).map(
      (item: RawVIndexItem): VIndexItem => ({
        ...item,
        normalizedFilename: item.filename.normalize('NFKC').toLowerCase(),
      }),
    );

    const sorted = vindex.sort((a: VIndexItem, b: VIndexItem) => {
      const dateDiff = b.date - a.date;
      return dateDiff == 0 ? a.filename.localeCompare(b.filename) : dateDiff;
    });

    const root = document.getElementById('root');
    render(React.createElement(App, { vindex: sorted }), root);
  }, 1000);
})();
