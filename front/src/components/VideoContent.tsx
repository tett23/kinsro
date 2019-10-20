import React from 'react';
import { useParams } from 'react-router';
import { useSelector } from 'react-redux';
import { basename } from 'path';
import { State } from 'modules';
import { VIndexItem } from 'models/vindex';

type OwnProps = {};

type StateProps = {
  digest: string;
  vindexItem: VIndexItem | null;
};

type DispatchProps = {};

export type VideoContentProps = OwnProps & StateProps & DispatchProps;

export function VideoContent({ digest, vindexItem }: VideoContentProps) {
  if (vindexItem == null) {
    return null;
  }

  return (
    <div>
      <h2>{basename(vindexItem.filename)}</h2>
      <video controls>
        <source src={`/videos/${digest}`} />
      </video>
    </div>
  );
}

export default function(ownProps: OwnProps) {
  return <VideoContent {...buildVideoContentProps(ownProps)} />;
}

export function buildVideoContentProps(ownProps: OwnProps): VideoContentProps {
  const { digest } = useParams<{ digest: string }>();
  const stateProps: StateProps = useSelector((state: State) => ({
    digest,
    vindexItem:
      state.domain.vindex.find((item) => item.digest === digest) || null,
  }));

  return {
    ...ownProps,
    ...stateProps,
  };
}
