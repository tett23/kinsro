import { randomInt } from 'crypto';

(async () => {
  async function bitonicSort<T>(
    up: boolean,
    arr: T[],
    compare: (a: T, b: T) => boolean,
  ): Promise<T[]> {
    if (arr.length === 1) {
      return arr;
    }

    const [a, b] = split2(arr);
    const ps = [
      await asyncJob(async () => await bitonicSort(true, a, compare)),
      await asyncJob(async () => await bitonicSort(false, b, compare)),
    ];

    return bitonicMerge(
      up,
      (await Promise.all(ps)).reduce((acc, item) => acc.concat(item), []),
      compare,
    );
  }

  function split2<T>(arr: T[]): [T[], T[]] {
    const mid = Math.floor(arr.length / 2);

    return [arr.slice(0, mid), arr.slice(mid)];
  }

  async function bitonicMerge<T>(
    up: boolean,
    arr: T[],
    compare: (a: T, b: T) => boolean,
  ): Promise<T[]> {
    if (arr.length === 1) {
      return arr;
    }

    const ps = split2(bitonicCompare(up, arr, compare)).map(
      async (item) => await asyncJob(() => bitonicMerge(up, item, compare)),
    );

    return (await Promise.all(ps)).reduce((acc, item) => acc.concat(item), []);
  }

  function bitonicCompare<T>(
    up: boolean,
    arr: T[],
    compare: (a: T, b: T) => boolean,
  ): T[] {
    const [a, b] = split2(arr);

    a.forEach((_, idx) => {
      if (compare(a[idx], b[idx]) === up) {
        [a[idx], b[idx]] = [b[idx], a[idx]];
      }
    });

    return a.concat(b);
  }

  function* seq() {
    let count = 0;
    while (true) {
      count += 1;

      yield count;
    }
  }

  const seqGen = seq();

  function requestIdleCallback(f: () => void): number {
    f();
    return seqGen.next().value as number;
  }

  async function asyncJob<T>(f: () => T): Promise<T> {
    return new Promise((resolve) => {
      requestIdleCallback(() => {
        resolve(f());
      });
    });
  }

  const arr = Array.from({ length: 2 ** 16 }).map(() => randomInt(2 ** 32));
  const compare = (a: number, b: number) => a > b;

  console.log(await bitonicSort(true, arr, compare)); // [ 1, 2, 3, 4, 5, 6, 10, 11]

  console.log(await bitonicSort(false, arr, compare)); // [11, 10, 6, 5, 4, 3, 2, 1]
})();
