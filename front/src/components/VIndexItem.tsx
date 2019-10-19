import React from 'react';
import { basename } from 'path';
import { Link } from 'react-router-dom';
import { VIndexItem as VIndexItemType } from 'models/VIndex';
import { FileDate } from './common/Date';

type OwnProps = {} & VIndexItemType;

type StateProps = {};

type DispatchProps = {};

export type VIndexItemProps = OwnProps & StateProps & DispatchProps;

export function VIndexItem({ digest, filename, date }: VIndexItemProps) {
  return (
    <tr>
      <td>
        <Link to={`/content/${digest}`}>{basename(filename)}</Link>
      </td>
      <td>
        <FileDate date={date} />
      </td>
    </tr>
  );
}

export default function VIndexItemDefault(ownProps: OwnProps) {
  return <VIndexItem {...buildVIndexItemProps(ownProps)} />;
}

export function buildVIndexItemProps(ownProps: OwnProps): VIndexItemProps {
  return {
    ...ownProps,
  };
}
