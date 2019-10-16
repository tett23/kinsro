import React from 'react';
import { VIndex as VIndexType } from 'models/VIndex';
import { useSelector } from 'react-redux';
import { State } from 'modules';
import VIndexItem from 'components/VIndexItem';

type OwnProps = {};

type StateProps = {
  vindex: VIndexType;
};

type DispatchProps = {};

export type VIndexProps = OwnProps & StateProps & DispatchProps;

export function VIndex({ vindex }: VIndexProps) {
  const items = vindex.map((item) => (
    <VIndexItem key={item.digest} {...item} />
  ));

  return (
    <table>
      <tbody>{items}</tbody>
    </table>
  );
}

export default function VIndexDefault() {
  return <VIndex {...buildVIndexProps()} />;
}

export function buildVIndexProps(): VIndexProps {
  const stateProps: StateProps = useSelector((state: State) => {
    const vindex = state.domain.vindex;
    const filterText = state.ui.filterText;
    if (filterText == null || filterText === '') {
      return { vindex };
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

    return {
      vindex: sorted,
    };
  });

  return {
    ...stateProps,
  };
}
