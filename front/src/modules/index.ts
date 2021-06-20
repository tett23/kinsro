import { VIndexItem } from 'models/vindex';
import { atom } from 'recoil';

export const mainVindexQuery = atom<string | null>({
  key: 'Main/FilterText',
  default: null,
});

export const mainVindex = atom<VIndexItem[]>({
  key: 'Main/Vindex',
  default: [],
});

export const mainCurrentPaginationPage = atom<number>({
  key: 'Main/PageCunt',
  default: 0,
});
