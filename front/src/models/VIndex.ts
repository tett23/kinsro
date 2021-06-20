export type VIndexItem = {
  storage: string;
  date: number;
  filename: string;
  digest: string;
};

function sortVIndex(vindex: VIndexItem[]): VIndexItem[] {
  return vindex.sort((a, b) => {
    const aTime = toTimestamp(a.filename, a.date);
    const bTime = toTimestamp(b.filename, b.date);

    if (aTime !== bTime) {
      return bTime - aTime;
    }

    return a.filename < b.filename ? -1 : 1;
  });
}

const DateTimeRe = /^(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})_.+/;
const DateRe = /^(\d{4})(\d{2})(\d{2})/;

function toTimestamp(filename: string, date: number): number {
  return Date.parse(
    DateTimeRe.test(filename)
      ? filename.replace(DateTimeRe, '$1-$2-$3T$4:$5:00')
      : `${date}`.replace(DateRe, '$1-$2-$3'),
  );
}

function filterVIndex(
  vindex: VIndexItem[],
  query: string | null,
): VIndexItem[] {
  if (query == null || query === '') {
    return vindex;
  }

  const normalizedQuery = query.normalize('NFKC').toLowerCase();

  return vindex.filter((item) => {
    return item.filename
      .normalize('NFKC')
      .toLowerCase()
      .includes(normalizedQuery);
  });
}

export function filterAndSort(
  vindex: VIndexItem[],
  filterText: string | null,
): VIndexItem[] {
  return sortVIndex(filterVIndex(vindex, filterText));
}
