const fs = require('fs');
const { join, basename, dirname } = require('path');

const Re = /^(?:-|rwx)+\s+\d+\s+tett23\s+staff\s+\d+\s+(\d+)\s+(\d+)\s+(\d+)\s+(.+)$/;
const StoragePath = '/Volumes/movie/frogbit';
const DstStoragePath = '/Volumes/video1';

function buildList() {
  const items = fs
    .readFileSync('./movies.txt', { encoding: 'utf8' })
    .split('\n')
    .slice(1)
    .filter((item) => Re.test(item))
    .map((item) => item.replace(Re, '$1\t$2\t$3\t$4').split('\t'))
    .map(([month, day, year, name]) => {
      return [
        name,
        `${year}/${month.padStart(2, 0)}/${day.padStart(2, 0)}/${name}`,
      ];
    });
  const json = JSON.stringify(items, null, 2);
  fs.writeFileSync('movies.json', json, { encoding: 'utf8' });
}

async function move() {
  const json = JSON.parse(
    fs.readFileSync('./movies.json', { encoding: 'utf8' }),
  );

  await copyFiles(json);
}

async function copyFiles(items) {
  while (items.length !== 0) {
    const item = items.shift();
    console.log(new Date(), items.length, item);
    if (item == null) {
      continue;
    }

    try {
      await copyFile(item);
    } catch (e) {
      console.log('error', new Date(), item, e);
      await addErrorItem(item);
    } finally {
      fs.writeFileSync('./movies.json', JSON.stringify(items, null, 2), {
        encoding: 'utf8',
      });
    }
  }
}

function copyFile([orig, toPath]) {
  const src = join(StoragePath, orig);
  const dst = join(DstStoragePath, toPath);
  // if (Math.floor(Math.random() * (10 - 0)) + 0 === 1) {
  //   throw new Error();
  // }
  fs.mkdirSync(dirname(dst), { recursive: true });
  fs.copyFileSync(src, dst);
}

async function addErrorItem(item) {
  const errorContents = await errorFileContents();

  errorContents.push(item);

  fs.writeFileSync(
    './movies_error.json',
    JSON.stringify(errorContents, null, 2),
  );
}

async function errorFileContents() {
  return (async () =>
    JSON.parse(
      fs.readFileSync('./movies_error.json', { encoding: 'utf8' }),
    ))().catch(() => []);
}

(async () => {
  buildList();
  await move();
})();

// const videoDirs = fs
//   .readdirSync(storagePath)
//   .filter((item) => /^\d{6}$/.test(item))
//   .map((item) => {
//     dirPath = join(storagePath, item);
//     return fs
//       .readdirSync(dirPath)
//       .filter((item) => /\.mp4$/.test(item))
//       .map((name) => join(dirPath, name));
//   })
//   .reduce((item, acc) => [...acc, ...item], [])
//   .map((item) => {
//     const date = item.split('/')[3].slice(0);
//     const year = parseInt(date.slice(0, 2)) + 2000;
//     const month = parseInt(date.slice(2, 4));
//     const day = parseInt(date.slice(4, 6));

//     return [[year, month, day], item];
//   })
//   .map(([[year, month, day], src]) => {
//     const dst = join(
//       storagePath,
//       year.toString(),
//       month.toString().padStart(2, '0'),
//       day.toString().padStart(2, '0'),
//       basename(src),
//     );

//     fs.mkdirSync(dirname(dst), { recursive: true });
//     fs.renameSync(src, dst);
//   });

// console.log(videoDirs);
