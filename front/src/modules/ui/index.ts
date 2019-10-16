export type UIState = {
  filterText: string | null;
};

export function initialUIState(): UIState {
  return {
    filterText: '',
  };
}

const UpdateFilterText = 'UI/UpdateFilterText';

export function updateFilterText(value: string) {
  return {
    type: UpdateFilterText,
    payload: value,
  };
}

export type UIActions = ReturnType<typeof updateFilterText>;

export default function uiReducer(
  state: UIState = initialUIState(),
  action: UIActions,
): UIState {
  switch (action.type) {
    case UpdateFilterText:
      return {
        ...state,
        filterText: action.payload,
      };
    default:
      return state;
  }
}
