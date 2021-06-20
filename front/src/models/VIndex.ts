export type RawVIndexItem = {
  storage: string;
  date: number;
  filename: string;
  digest: string;
};

export type VIndexItem = RawVIndexItem & {
  normalizedFilename: string;
  timestamp: number;
};

export function sortVIndex(vindex: VIndexItem[]): VIndexItem[] {
  return vindex.sort((a, b) => {
    if (a.timestamp !== b.timestamp) {
      return b.timestamp - a.timestamp;
    }

    return a.filename < b.filename ? -1 : 1;
  });
}

const DateTimeRe = /^(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})_.+/;
const DateRe = /^(\d{4})(\d{2})(\d{2})/;

export function toTimestamp(filename: string, date: number): number {
  return Date.parse(
    DateTimeRe.test(filename)
      ? filename.replace(DateTimeRe, '$1-$2-$3T$4:$5:00')
      : `${date}`.replace(DateRe, '$1-$2-$3'),
  );
}
