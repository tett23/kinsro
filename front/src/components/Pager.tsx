import React from 'react';

type OwnProps = {
  current: number;
  max: number;
  onClickPage: (page: number) => void;
};

type StateProps = {};

type DispatchProps = {};

export type PagerProps = OwnProps & StateProps & DispatchProps;

export function Pager({ current, max, onClickPage }: PagerProps) {
  const items = Array.from({ length: max }).map((_, i) => (
    <PageItem
      key={`page-item-${i}`}
      page={i}
      isCurrent={current === i}
      onClick={() => onClickPage(i)}
    />
  ));

  return <ul>{items}</ul>;
}

export default function(ownProps: OwnProps) {
  return <Pager {...buildPagerProps(ownProps)} />;
}

export function buildPagerProps(ownProps: OwnProps): PagerProps {
  return {
    ...ownProps,
  };
}

type PageItemProps = {
  page: number;
  isCurrent: boolean;
  onClick: () => void;
};

function PageItem({ page, onClick }: PageItemProps) {
  return <li onClick={onClick}>{page}</li>;
}
