export type UIState = {
  filterText: string | null;
  pages: {
    root: {
      page: number | null;
    };
  };
};

export function initialUIState(): UIState {
  return {
    filterText: '',
    pages: {
      root: {
        page: 0,
      },
    },
  };
}

const UpdateFilterText = 'UI/UpdateFilterText' as const;

export function updateFilterText(value: string) {
  return {
    type: UpdateFilterText,
    payload: value,
  };
}

const RootUpdatePage = 'UI/Root/UpdatePage' as const;

export function rootUpdatePage(value: number) {
  return {
    type: RootUpdatePage,
    payload: value,
  };
}

export type UIActions =
  | ReturnType<typeof updateFilterText>
  | ReturnType<typeof rootUpdatePage>;

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
    case RootUpdatePage:
      return {
        ...state,
        pages: {
          ...state.pages,
          root: {
            ...state.pages.root,
            page: action.payload,
          },
        },
      };
    default:
      return state;
  }
}
