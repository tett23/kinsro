import { asyncJob } from './dummyRequestIdleCallback';

export const asyncBitonicSort = bitonicSort;

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
    await asyncSort(true, a, compare),
    await asyncSort(false, b, compare),
  ];

  return bitonicMerge(
    up,
    (await Promise.all(ps)).reduce((acc, item) => acc.concat(item), []),
    compare,
  );
}

const asyncSort = asyncJob(bitonicSort);

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
    async (item) => await asyncMerge(up, item, compare),
  );

  return (await Promise.all(ps)).reduce((acc, item) => acc.concat(item), []);
}

const asyncMerge = asyncJob(bitonicMerge);

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
