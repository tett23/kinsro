import React from 'react';
import { VIndex as VIndexType } from 'models/VIndex';
import { useSelector, useDispatch } from 'react-redux';
import { State } from 'modules';
import VIndexItem from 'components/VIndexItem';
import { Pager } from './Pager';
import { rootUpdatePage } from 'modules/ui';

type OwnProps = {
  page: number;
};

type StateProps = {
  vindex: VIndexType;
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
  const stateProps: StateProps = useSelector((state: State) => {
    const vindex = filterAndSort(state.domain.vindex, state.ui.filterText);
    const page = ownProps.page;
    const maxPages = Math.ceil(vindex.length / 100);

    return {
      vindex: vindex.slice(page * 100, (page + 1) * 100),
      maxPages,
    };
  });

  const dispatch = useDispatch();
  const dispatchProps: DispatchProps = {
    onClickPage: (page: number) => {
      dispatch(rootUpdatePage(page));
    },
  };

  return {
    ...ownProps,
    ...stateProps,
    ...dispatchProps,
  };
}

function filterAndSort(vindex: VIndex, filterText: string | null): VIndex {
  if (filterText == null || filterText === '') {
    return vindex;
  }

  const normalizedQuery = filterText.normalize('NFKC');
  const filtered = vindex.filter((item) => {
    return item.filename.normalize('NFKC').includes(normalizedQuery);
  });
  const sorted = filtered.sort((a, b) => {
    if (a.date !== b.date) {
      return b.date - a.date;
    }

    return a.filename < b.filename ? -1 : 1;
  });

  return sorted;
}
