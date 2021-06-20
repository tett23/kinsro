import React from 'react';
import FileFilter from './FileFilter';
import VIndex from './VIndex';
import { useRecoilValue } from 'recoil';
import { mainCurrentPaginationPage } from 'modules';

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

export default function (ownProps: OwnProps) {
  return <Root {...buildRootProps(ownProps)} />;
}

export function buildRootProps(_: OwnProps): RootProps {
  const page = useRecoilValue(mainCurrentPaginationPage);

  return {
    page,
  };
}
