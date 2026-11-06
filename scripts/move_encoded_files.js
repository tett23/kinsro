const fs = require('fs');
const { join, basename, dirname, extname } = require('path');

const StoragePath = '/media/video1';
const TmpPath = '/media/video_tmp';
const groups = fs
  .readdirSync(TmpPath)
  .map((item) => {
    return item;
  })
  .reduce((acc, item) => {
    const base = basename(item);
    const name = base.slice(0, base.lastIndexOf(extname(base)));

    if (acc[name] == null) {
      acc[name] = [];
    }

    acc[name].push(item);

    return acc;
  }, {});

Object.entries(groups)
  .reduce((acc, [name, items]) => {
    const isIncludeTS = items.some((item) => item.endsWith('.ts'));
    if (isIncludeTS) {
      return acc;
    }
    if (!name.match(/^\d{8}/)) {
      return acc;
    }

    acc.push(...items);

    return acc;
  }, [])
  .map((item) => {
    return [
      join(TmpPath, item),
      join(
        StoragePath,
        item.slice(0, 4),
        item.slice(4, 6),
        item.slice(6, 8),
        item,
      ),
    ];
  })
  .forEach(([src, dst], i) => {
    console.log(new Date(), i, src, dst);
    try {
      fs.copyFileSync(src, dst);
      fs.unlinkSync(src);
    } catch (e) {
      console.log(e);
    }
  });
