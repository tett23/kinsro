import 'core-js/stable';
import 'regenerator-runtime/runtime';
import './wasm_exec';
import React from 'react';
import { render } from 'react-dom';
import App from './App';
import {
  RawVIndexItem,
  sortVIndex,
  toTimestamp,
  VIndexItem,
} from 'models/vindex';

(async () => {
  const go = new global.Go();
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
    console.time('a');
    global.parseVIndex(bytes);
    console.timeEnd('a');

    const vindex = build();

    console.time('sort');
    const sorted = await sortVIndex(vindex);
    console.timeEnd('sort');

    const root = document.getElementById('root');
    render(React.createElement(App, { vindex: sorted }), root);
  }, 1);
})();

function build(): VIndexItem[] {
  console.time('parseJSON');
  const rawVIndex: RawVIndexItem[] = JSON.parse(global.vindex);
  console.timeEnd('parseJSON');

  console.time('b');
  rawVIndex.forEach((item: RawVIndexItem): VIndexItem => {
    (item as VIndexItem).normalizedFilename = item.filename
      .normalize('NFKC')
      .toLowerCase();
    (item as VIndexItem).timestamp = toTimestamp(item.filename, item.date);

    return item as VIndexItem;
  });
  console.timeEnd('b');

  return rawVIndex as VIndexItem[];
}
