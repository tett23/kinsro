import 'core-js/stable';
import 'regenerator-runtime/runtime';
import './wasm_exec';

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
  console.log(result);
  console.log(result.module);

  go.run(result.instance);
})();
