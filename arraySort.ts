import { randomInt } from 'crypto';

(() => {
  const arr = Array.from({ length: 2 ** 16 }).map(() => randomInt(2 ** 32));
  const compare = (a: number, b: number) => a - b;

  console.time('sort');
  console.log(arr.sort(compare));
  console.timeEnd('sort');
})();
