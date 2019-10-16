import React from 'react';

type OwnProps = {
  date: number;
};

type StateProps = {};

type DispatchProps = {};

export type FileDateProps = OwnProps & StateProps & DispatchProps;

export function FileDate({ date }: FileDateProps) {
  const characters = date.toString().split('');
  const year = characters.slice(0, 4);
  const month = characters.slice(4, 6);
  const day = characters.slice(6, 8);

  return (
    <>
      {year}/{month}/{day}
    </>
  );
}

export default function(ownProps: OwnProps) {
  return <FileDate {...buildFileDateProps(ownProps)} />;
}

export function buildFileDateProps(ownProps: OwnProps): FileDateProps {
  return {
    ...ownProps,
  };
}
