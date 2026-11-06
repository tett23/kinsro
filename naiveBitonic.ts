import { randomInt } from 'crypto';

(() => {
  function bitonicSort<T>(
    up: boolean,
    arr: T[],
    compare: (a: T, b: T) => boolean,
  ): T[] {
    if (arr.length === 1) {
      return arr;
    }

    const [a, b] = split2(arr);

    return bitonicMerge(
      up,
      bitonicSort(true, a, compare).concat(bitonicSort(false, b, compare)),
      compare,
    );
  }

  function split2<T>(arr: T[]): [T[], T[]] {
    const mid = Math.floor(arr.length / 2);

    return [arr.slice(0, mid), arr.slice(mid)];
  }

  function bitonicMerge<T>(
    up: boolean,
    arr: T[],
    compare: (a: T, b: T) => boolean,
  ): T[] {
    if (arr.length === 1) {
      return arr;
    }

    const [a, b] = split2(bitonicCompare(up, arr, compare));

    return bitonicMerge(up, a, compare).concat(bitonicMerge(up, b, compare));
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

  const arr = Array.from({ length: 2 ** 3 }).map(() => randomInt(2 ** 32));
  const compare = (a: number, b: number) => a > b;

  console.time('sort');
  console.log(bitonicSort(true, arr, compare)); // [ 1, 2, 3, 4, 5, 6, 10, 11]
  console.timeEnd('sort');

  console.log(bitonicSort(false, arr, compare)); // [11, 10, 6, 5, 4, 3, 2, 1]
})();
