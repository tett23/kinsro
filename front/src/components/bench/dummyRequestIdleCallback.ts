// function* seq() {
//   let count = 0;
//   while (true) {
//     count += 1;

//     yield count;
//   }
// }

// const seqGen = seq();

// function requestIdleCallback(f: () => void): number {
//   f();
//   return seqGen.next().value as number;
// }

type UnwrapPromise<T> = T extends Promise<infer U> ? U : T;

export function asyncJob<R>(f: () => R): () => Promise<UnwrapPromise<R>>;
export function asyncJob<T1, R>(
  f: (arg1: T1) => R,
): (arg1: T1) => Promise<UnwrapPromise<R>>;
export function asyncJob<T1, T2, R>(
  f: (arg1: T1, arg2: T2) => R,
): (arg1: T1, arg2: T2) => Promise<UnwrapPromise<R>>;
export function asyncJob<T1, T2, T3, R>(
  f: (arg1: T1, arg2: T2, arg3: T3) => R,
): (arg1: T1, arg2: T2, arg3: T3) => Promise<UnwrapPromise<R>>;

export function asyncJob(f: any): any {
  return (...args: any) =>
    new Promise((resolve) => {
      global.requestIdleCallback(() => {
        const ret = f(...args);
        if (ret instanceof Promise) {
          ret.then(resolve);
          return;
        }

        resolve(ret);
      });
    });
}
