import { bitonicSortWithPadding } from 'utils/asyncSort';

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

export async function sortVIndex(vindex: VIndexItem[]): Promise<VIndexItem[]> {
  return bitonicSortWithPadding(false, vindex, (a, b) => {
    if (a.timestamp !== b.timestamp) {
      return a.timestamp > b.timestamp;
    }

    return a.filename > b.filename;
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
