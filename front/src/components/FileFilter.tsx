import React from 'react';
import { mainVindexQuery } from 'modules';
import { useRecoilState } from 'recoil';

export type FileFilterProps = {
  filterText: string | null;
  onChange: (value: string | null) => void;
};

export function FileFilter({ filterText, onChange }: FileFilterProps) {
  return (
    <div>
      <input
        type="text"
        defaultValue={filterText ?? ''}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
}

export default function () {
  return <FileFilter {...buildFileFilterProps()} />;
}

export function buildFileFilterProps(): FileFilterProps {
  const [filterText, setFilterText] = useRecoilState(mainVindexQuery);

  return {
    filterText,
    onChange: (value: string | null) =>
      setFilterText(value === '' ? null : value),
  };
}
