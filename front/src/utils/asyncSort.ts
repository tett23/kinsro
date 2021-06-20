import { asyncJob } from 'utils/promiseRequestIdleCallback';

class PaddingItem {}
const PaddingItemInstance = new PaddingItem();

const Threshold = 2 ** 12;

function setPad<T>(arr: T[]): Array<T | PaddingItem> {
  const f = (v: number, c: number): number =>
    v === 0 ? c : f(Math.floor(v / 2), c + 1);

  return Array.from({ length: 2 ** f(arr.length, 0) }).map((_, idx) =>
    arr[idx] == null ? PaddingItemInstance : arr[idx],
  );
}

export async function bitonicSortWithPadding<T>(
  up: boolean,
  arr: T[],
  compare: (a: T, b: T) => boolean,
): Promise<T[]> {
  return (await bitonicSortFn(up, setPad(arr), compare)).filter(
    (item): item is T => !(item instanceof PaddingItem),
  );
}

function bitonicSortFn<T>(
  up: boolean,
  arr: Array<T | PaddingItem>,
  compare: (a: T, b: T) => boolean,
): Array<T | PaddingItem> | Promise<Array<T | PaddingItem>> {
  return arr.length > Threshold
    ? asyncBitonicSort(up, arr, compare).then((ret) => ret)
    : syncBitonicSort(up, arr, compare);
}

function syncBitonicSort<T>(
  up: boolean,
  arr: Array<T | PaddingItem>,
  compare: (a: T, b: T) => boolean,
): Array<T | PaddingItem> | Promise<Array<T | PaddingItem>> {
  if (arr.length === 1) {
    return arr;
  }

  const [a, b] = split2(arr);
  const items = [
    bitonicSortFn(true, a, compare),
    bitonicSortFn(false, b, compare),
  ];
  if (!(items[0] instanceof Promise)) {
    const result = (items as any as Array<T | PaddingItem>[]).reduce(
      (acc, item) => acc.concat(item),
      [],
    );
    return mergeFn(up, result, compare);
  }

  return Promise.all(items)
    .then((result) => result.reduce((acc, item) => acc.concat(item), []))
    .then((result) => {
      return mergeFn(up, result, compare);
    });
}

const asyncBitonicSort = asyncJob(syncBitonicSort);

function split2<T>(arr: T[]): [T[], T[]] {
  const mid = Math.floor(arr.length / 2);

  return [arr.slice(0, mid), arr.slice(mid)];
}

function mergeFn<T>(
  up: boolean,
  arr: Array<T | PaddingItem>,
  compare: (a: T, b: T) => boolean,
): Array<T | PaddingItem> | Promise<Array<T | PaddingItem>> {
  return arr.length > Threshold
    ? asyncMerge(up, arr, compare).then((ret) => ret)
    : syncMerge(up, arr, compare);
}

function syncMerge<T>(
  up: boolean,
  arr: Array<T | PaddingItem>,
  compare: (a: T, b: T) => boolean,
): Array<T | PaddingItem> | Promise<Array<T | PaddingItem>> {
  if (arr.length === 1) {
    return arr;
  }

  const [a, b] = split2(bitonicCompare(up, arr, compare));
  const items = [mergeFn(up, a, compare), mergeFn(up, b, compare)];
  if (!(items[0] instanceof Promise)) {
    return (items as any as Array<T | PaddingItem>[]).reduce(
      (acc, item) => acc.concat(item),
      [],
    );
  }

  return Promise.all(items).then((ret) =>
    ret.reduce((acc, item) => acc.concat(item), []),
  );
}

const asyncMerge = asyncJob(syncMerge);

function bitonicCompare<T>(
  up: boolean,
  arr: Array<T | PaddingItem>,
  compare: (a: T, b: T) => boolean,
): Array<T | PaddingItem> {
  const [a, b] = split2(arr);

  a.forEach((_, idx) => {
    if (a[idx] instanceof PaddingItem) {
      [a[idx], b[idx]] = [b[idx], a[idx]];
      return;
    }
    if (!up && b[idx] instanceof PaddingItem) {
      [a[idx], b[idx]] = [b[idx], a[idx]];
      return;
    }

    if (compare(a[idx] as T, b[idx] as T) === up) {
      [a[idx], b[idx]] = [b[idx], a[idx]];
    }
  });

  return a.concat(b);
}
