// import { randomInt } from 'crypto';
// import { asyncBitonicSort } from './asyncBitonic';
// import { bitonicSortWithPadding } from './asyncBitonic2';
import { bitonicSortWithPadding } from 'utils/asyncSort';
import { naiveBitonicSort } from './naiveBitonic';

export async function bench() {
  const arr = [new Uint16Array(2 ** 15), new Uint16Array(2 ** 15)].reduce(
    // const arr = [new Uint16Array(2 ** 10)].reduce(
    (acc, a) => {
      window.crypto.getRandomValues(a);
      a.forEach((v) => acc.push(v));

      return acc;
    },
    [] as Array<number>,
  );
  console.log(arr.length);
  // const arr = Array.from({ length: 2 ** 16 }).map(() => randomInt(2 ** 32));
  const compare = (a: number, b: number) => a > b;

  const copy = arr.reduce((acc, item) => {
    acc.push(item);
    return acc;
  }, [] as number[]);
  console.time('Array.sort');
  copy.sort();
  console.timeEnd('Array.sort');

  console.time('naive');
  naiveBitonicSort(true, arr, compare);
  console.timeEnd('naive');

  console.time('async2');
  // await bitonicSortWithPadding(true, arr, compare);
  await bitonicSortWithPadding(
    true,
    arr.slice(0, arr.length - 1),
    (a, b) => a > b,
  );
  console.timeEnd('async2');

  // console.time('async');
  // await asyncBitonicSort(true, arr, compare);
  // console.timeEnd('async');
}
