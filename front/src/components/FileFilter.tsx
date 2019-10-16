import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { State } from 'modules';
import { updateFilterText } from 'modules/ui';

type OwnProps = {};

type StateProps = {
  filterText: string | null;
};

type DispatchProps = {
  onChange: (value: string) => void;
};

export type FileFilterProps = OwnProps & StateProps & DispatchProps;

export function FileFilter({ filterText, onChange }: FileFilterProps) {
  return (
    <div>
      <input
        type="text"
        defaultValue={filterText || ''}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
}

export default function(ownProps: OwnProps) {
  return <FileFilter {...buildFileFilterProps(ownProps)} />;
}

export function buildFileFilterProps(ownProps: OwnProps): FileFilterProps {
  const stateProps: StateProps = useSelector((state: State) => ({
    filterText: state.ui.filterText,
  }));
  const dispatch = useDispatch();
  const dispatchProps: DispatchProps = {
    onChange: (value) => {
      dispatch(updateFilterText(value));
    },
  };

  return {
    ...ownProps,
    ...stateProps,
    ...dispatchProps,
  };
}
