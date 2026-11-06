import React, { useState } from 'react';
import { useParams } from 'react-router';
import { basename } from 'path';
import { VIndexItem } from 'models/vindex';
import useKVS from 'kvs/hooks';
import { useRecoilValue } from 'recoil';
import { mainVindex } from 'modules';
import { Link } from 'react-router-dom';

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
  const [playing, setPlaying] = useState(false);

  if (vindexItem == null) {
    return null;
  }

  return (
    <figure id="player">
      {!playing && (
        <div className="caption">
          <span className="back">
            <Link to="/" role="button">
              &lt;
            </Link>
          </span>
          <h1>{basename(vindexItem.filename)}</h1>
        </div>
      )}
      <video
        controls
        ref={(ref: HTMLVideoElement | null) => {
          if (ref != null) {
            ref.volume = volume;
          }
        }}
        onPlay={() => {
          setPlaying(true);
        }}
        onPause={() => {
          setPlaying(false);
        }}
        onVolumeChange={(e) =>
          onVolumeChange((e.target as HTMLVideoElement).volume)
        }
      >
        <source src={`/videos/${digest}.mp4`} />
        {/* <source src="http://10.0.1.110/videos/9ac2a341946632330f872bcffca8eb88.mp4"></source> */}
      </video>
    </figure>
  );
}

export default function (ownProps: OwnProps) {
  return <VideoContent {...buildVideoContentProps(ownProps)} />;
}

export function buildVideoContentProps(ownProps: OwnProps): VideoContentProps {
  const vindex = useRecoilValue(mainVindex);
  const { digest } = useParams<{ digest: string }>();
  const [volume, setVolume] = useKVS<number>('volume', 1.0);

  const stateProps = {
    digest,
    vindexItem: vindex.find((item) => item.digest === digest) || null,
    volume: volume ?? 1.0,
  };
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
