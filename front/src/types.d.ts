namespace NodeJS {
  namespace WebAssembly {
    function instantiateStreaming(fetch: Promise<Response>, obj: any): any;
  }

  interface IdleDeadline {
    timeRemaining(): DOMHighResTimeStamp;
    readonly didTimeout: boolean;
  }

  interface IdleRequestOptions {
    timeout: number;
  }

  type IdleCallbackHandle = number;

  type IdleRequestCallback = (deadline: IdleDeadline) => void;

  interface Global {
    Go: any;
    vindex: string;
    parseVIndex: (buf: Uint8Array) => void;

    requestIdleCallback: (
      callback: IdleRequestCallback,
      options?: IdleRequestOptions,
    ) => number;
    cancelIdleCallback: (handle: number) => void;
  }
}

declare module 'utils/promiseRequestIdleCallback' {
  function asyncJob<R>(f: () => R): () => Promise<R>;
  function asyncJob<T1, R>(f: (arg1: T1) => R): (arg1: T1) => Promise<R>;
  function asyncJob<T1, T2, R>(
    f: (arg1: T1, arg2: T2) => R,
  ): (arg1: T1, arg2: T2) => Promise<R>;
  function asyncJob<T1, T2, T3, R>(
    f: (arg1: T1, arg2: T2, arg3: T3) => R,
  ): (arg1: T1, arg2: T2, arg3: T3) => Promise<R>;
}
