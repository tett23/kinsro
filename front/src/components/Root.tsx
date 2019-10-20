import React from 'react';
import FileFilter from './FileFilter';
import VIndex from './VIndex';
import { useSelector } from 'react-redux';
import { State } from 'modules';

type OwnProps = {};

type StateProps = {
  page: number;
};

type DispatchProps = {};

export type RootProps = OwnProps & StateProps & DispatchProps;

export function Root({ page }: RootProps) {
  return (
    <div>
      <FileFilter />
      <VIndex page={page} />
    </div>
  );
}

export default function(ownProps: OwnProps) {
  return <Root {...buildRootProps(ownProps)} />;
}

export function buildRootProps(ownProps: OwnProps): RootProps {
  const stateProps: StateProps = useSelector((state: State) => ({
    page: state.ui.pages.root.page || 0,
  }));

  return {
    ...ownProps,
    ...stateProps,
  };
}
