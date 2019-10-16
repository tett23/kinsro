import { VIndex } from 'models/vindex';
import { AnyAction } from 'redux';

export type DomainState = {
  vindex: VIndex;
};

export function initialDomainState(): DomainState {
  return {
    vindex: [],
  };
}

export type DomainActions = AnyAction;

export default function domainReducer(
  state: DomainState = initialDomainState(),
  action: DomainActions,
): DomainState {
  switch (action.type) {
    default:
      return state;
  }
}
