import { AnyAction } from 'redux';

export type UIState = {};

export function initialUIState(): UIState {
  return {};
}

export type UIActions = AnyAction;

export default function uiReducer(
  state: UIState = initialUIState(),
  action: UIActions,
): UIState {
  switch (action.type) {
    default:
      return state;
  }
}
