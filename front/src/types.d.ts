export {};

declare global {
  const Go: any;
  namespace WebAssembly {
    function instantiateStreaming(fetch: Promise<Response>, obj: any): any;
  }

  interface Window {
    vindex: string;
    parseVIndex: (buf: Uint8Array) => void;
  }
}
