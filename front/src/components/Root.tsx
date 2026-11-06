import React, { useEffect, useState } from 'react';
import FileFilter from './FileFilter';
import VIndex from './VIndex';
import { useRecoilState, useRecoilValue } from 'recoil';
import { mainCurrentPaginationPage, mainVindex } from 'modules';
import { SortDirection, sortVIndex } from 'models/VIndex';
import { Link } from 'react-router-dom';

type OwnProps = {};

type StateProps = {
  page: number;
};

type DispatchProps = {};

export type RootProps = OwnProps & StateProps & DispatchProps;

export function Root({ page }: RootProps) {
  return (
    <div>
      <h1>
        <Link to="/">kinsro</Link>
      </h1>
      <div>
        <FileFilter />
        {/* <Reverse /> */}
      </div>
      {/* <Counter /> */}
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

function Counter() {
  const [count, setCount] = useState(0);
  useEffect(() => {
    setTimeout(() => {
      setCount(count + 1);
    }, 1000);
  }, [count]);

  return <div>{count}</div>;
}

function Reverse() {
  const [vindex, setVindex] = useRecoilState(mainVindex);
  const [direction, setDirection] = useState<SortDirection>('desc');
  const onClick = () => {
    sortVIndex(vindex, direction).then((result) => setVindex(result));
    setDirection(direction === 'asc' ? 'desc' : 'asc');
  };

  return (
    <button onClick={onClick}>{direction === 'asc' ? '降順' : '昇順'}</button>
  );
}
