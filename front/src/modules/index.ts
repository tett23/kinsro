import { applyMiddleware, combineReducers, createStore, Store } from 'redux';
import thunk, { ThunkDispatch, ThunkMiddleware } from 'redux-thunk';
import domainReducer, {
  DomainActions,
  DomainState,
  initialDomainState,
} from 'modules/domain';
import uiReducer, { UIActions, UIState, initialUIState } from 'modules/ui';
import { VIndex } from 'models/vindex';

export type State = {
  domain: DomainState;
  ui: UIState;
};

export type Actions = DomainActions | UIActions;

export const reducers = combineReducers<State>({
  domain: domainReducer,
  ui: uiReducer,
});

export function initialState(): State {
  return {
    domain: initialDomainState(),
    ui: initialUIState(),
  };
}

export type ThunkStore<S, A extends Actions> = Store<S, A> & {
  dispatch: ThunkDispatch<S, undefined, A>;
};

export default function initializeStore(vindex: VIndex): Store<State, Actions> {
  const init: State = {
    domain: { vindex },
    ui: {},
  };

  const store: Store<State, Actions> = createStore(
    reducers as any,
    init,
    applyMiddleware(thunk as ThunkMiddleware<State, Actions>),
  );

  return store;
}
