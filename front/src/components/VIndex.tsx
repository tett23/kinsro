import React from 'react';
import { VIndexItem as VIndexItemType } from 'models/VIndex';
import VIndexItem from 'components/VIndexItem';
import { Pager } from './Pager';
import { useRecoilValue, useSetRecoilState } from 'recoil';
import {
  mainCurrentPaginationPage,
  mainVindex,
  mainVindexQuery,
} from 'modules';

type OwnProps = {
  page: number;
};

type StateProps = {
  vindex: VIndexItemType[];
  maxPages: number;
};

type DispatchProps = {
  onClickPage: (page: number) => void;
};

export type VIndexProps = OwnProps & StateProps & DispatchProps;

export function VIndex({ vindex, page, maxPages, onClickPage }: VIndexProps) {
  const items = vindex.map((item) => (
    <VIndexItem key={item.digest} {...item} />
  ));

  return (
    <div>
      <table>
        <tbody>{items}</tbody>
      </table>
      <Pager current={page} max={maxPages} onClickPage={onClickPage} />
    </div>
  );
}

export default function VIndexDefault(ownProps: OwnProps) {
  return <VIndex {...buildVIndexProps(ownProps)} />;
}

export function buildVIndexProps(ownProps: OwnProps): VIndexProps {
  const vindex = useRecoilValue(mainVindex);
  const filterText = useRecoilValue(mainVindexQuery);
  const setPage = useSetRecoilState(mainCurrentPaginationPage);
  const filtered = filterVIndex(vindex, filterText);
  const page = ownProps.page;
  const maxPages = Math.ceil(vindex.length / 100);

  const stateProps: StateProps = {
    vindex: filtered.slice(page * 100, (page + 1) * 100),
    maxPages,
  };

  const dispatchProps: DispatchProps = {
    onClickPage: (page: number) => {
      setPage(page);
    },
  };

  return {
    ...ownProps,
    ...stateProps,
    ...dispatchProps,
  };
}

function filterVIndex(
  vindex: VIndexItemType[],
  query: string | null,
): VIndexItemType[] {
  if (query == null || query === '') {
    return vindex;
  }

  const normalizedQuery = query.normalize('NFKC').toLowerCase();

  return vindex.filter((item) => {
    return item.normalizedFilename.includes(normalizedQuery);
  });
}
