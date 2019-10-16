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
  const stateProps: StateProps = useSelector((state: State) => ({
    vindex: state.domain.vindex,
  }));

  return {
    ...stateProps,
  };
}
