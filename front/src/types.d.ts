export {};

declare global {
  const Go: any;
  namespace WebAssembly {
    function instantiateStreaming(fetch: Promise<Response>, obj: any): any;
  }

  interface Window {
    vindex: string;
  }
}
