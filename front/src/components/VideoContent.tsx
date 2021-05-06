import React from 'react';
import { useParams } from 'react-router';
import { useSelector } from 'react-redux';
import { basename } from 'path';
import { State } from 'modules';
import { VIndexItem } from 'models/vindex';
import useKVS from 'kvs/hooks';

type OwnProps = {};

type StateProps = {
  digest: string;
  vindexItem: VIndexItem | null;
  volume: number;
};

type DispatchProps = {
  onVolumeChange: (value: number) => void;
};

export type VideoContentProps = OwnProps & StateProps & DispatchProps;

export function VideoContent({
  digest,
  vindexItem,
  volume,
  onVolumeChange,
}: VideoContentProps) {
  if (vindexItem == null) {
    return null;
  }

  return (
    <div>
      <h2>{basename(vindexItem.filename)}</h2>
      <video
        controls
        ref={(ref) => {
          if (ref != null) {
            ref.volume = volume;
          }
        }}
        onVolumeChange={(e) =>
          onVolumeChange((e.target as HTMLVideoElement).volume)
        }
        style={{ width: '100%', outline: 'none' }}
      >
        <source src={`/videos/${digest}.mp4`} />
      </video>
    </div>
  );
}

export default function(ownProps: OwnProps) {
  return <VideoContent {...buildVideoContentProps(ownProps)} />;
}

export function buildVideoContentProps(ownProps: OwnProps): VideoContentProps {
  const { digest } = useParams<{ digest: string }>();
  const [volume, setVolume] = useKVS<number>('volume', 1.0);

  const stateProps: StateProps = useSelector((state: State) => ({
    digest,
    vindexItem:
      state.domain.vindex.find((item) => item.digest === digest) || null,
    volume: volume ?? 1.0,
  }));
  const dispatchProps: DispatchProps = {
    onVolumeChange: (value: number) => {
      setVolume(value);
    },
  };

  return {
    ...ownProps,
    ...stateProps,
    ...dispatchProps,
  };
}
