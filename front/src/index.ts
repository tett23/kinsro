import 'core-js/stable';
import 'regenerator-runtime/runtime';
import './wasm_exec';
import React from 'react';
import { render } from 'react-dom';
import App from './App';

(async () => {
  const go = new Go();
  const result = await WebAssembly.instantiateStreaming(
    fetch('./kinsro.wasm'),
    go.importObject,
  ).catch((err: Error) => err);
  if (result instanceof Error) {
    console.error(result);
    return;
  }

  go.run(result.instance);
  const vindex = JSON.parse(window.vindex);

  console.log(window.vf());
  const root = document.getElementById('root');
  render(React.createElement(App, { vindex }), root);
})();
