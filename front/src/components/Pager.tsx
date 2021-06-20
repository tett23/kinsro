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

  return <ul style={{ listStyleType: 'none', padding: 0 }}>{items}</ul>;
}

export default function (ownProps: OwnProps) {
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

function PageItem({ page, isCurrent, onClick }: PageItemProps) {
  return (
    <li
      onClick={onClick}
      role="button"
      style={{
        display: 'inline-block',
        marginRight: '1rem',
        cursor: 'pointer',
        color: isCurrent ? 'black' : 'blue',
      }}
    >
      {page}
    </li>
  );
}
