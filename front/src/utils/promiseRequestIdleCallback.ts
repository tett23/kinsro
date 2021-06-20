export function asyncJob<R>(
  f: (...args: any[]) => R,
): (...args: any[]) => Promise<R> {
  return (...args: any[]) =>
    new Promise((resolve) => {
      global.requestIdleCallback(() => {
        resolve(f(...args));
      });
    });
}
