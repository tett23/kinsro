import 'core-js/stable';
import 'regenerator-runtime/runtime';
import './wasm_exec';
import React from 'react';
import { render } from 'react-dom';
import App from './App';
import { RawVIndexItem, toTimestamp, VIndexItem } from 'models/VIndex';
import { bench } from './components/bench';

(async () => {
  await bench();
  const go = new global.Go();
  const buf = await fetch(new URL('/kinsro.wasm', location.origin)).then(
    (resp) => resp.arrayBuffer(),
  );
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
    console.time('parseVIndex');
    global.parseVIndex(bytes);
    console.timeEnd('parseVIndex');
    console.time('parseJSON');
    const vindex = buildVIndex();
    console.timeEnd('parseJSON');

    const root = document.getElementById('root');
    render(React.createElement(App, { vindex }), root);
  }, 1);
})();

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
