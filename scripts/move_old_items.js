const fs = require('fs');
const { join, basename, dirname } = require('path');

const storagePath = '/Volumes/video1';

const videoDirs = fs
  .readdirSync(storagePath)
  .filter((item) => /^\d{6}$/.test(item))
  .map((item) => {
    dirPath = join(storagePath, item);
    return fs
      .readdirSync(dirPath)
      .filter((item) => /\.mp4$/.test(item))
      .map((name) => join(dirPath, name));
  })
  .reduce((item, acc) => [...acc, ...item], [])
  .map((item) => {
    const date = item.split('/')[3].slice(0);
    const year = parseInt(date.slice(0, 2)) + 2000;
    const month = parseInt(date.slice(2, 4));
    const day = parseInt(date.slice(4, 6));

    return [[year, month, day], item];
  })
  .map(([[year, month, day], src]) => {
    const dst = join(
      storagePath,
      year.toString(),
      month.toString().padStart(2, '0'),
      day.toString().padStart(2, '0'),
      basename(src),
    );

    fs.mkdirSync(dirname(dst), { recursive: true });
    fs.renameSync(src, dst);
  });

console.log(videoDirs);
