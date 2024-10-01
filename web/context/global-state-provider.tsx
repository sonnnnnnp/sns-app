"use client";

import { components } from "@/lib/api/client";
import { Cookie } from "@/lib/cookies";
import { GlobalState } from "@/types/global-state";
import { createContext, useState } from "react";

const initialState: GlobalState = {
  session: {
    user: null,
  },
  cookie: {
    accessToken: null,
    refreshToken: null,
  },
};

export const GlobalStateContext = createContext<{
  state: GlobalState;
  updateSessionUser: (user: components["schemas"]["User"]) => void;
}>({
  state: initialState,
  updateSessionUser: () => {},
});

export const GlobalStateProvider = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const [state, setState] = useState<GlobalState>(() => {
    // TODO: refactor
    const state = initialState;

    try {
      const user = JSON.parse(localStorage.getItem("user")!);
      state.session = {
        user: user ? user : null,
      };
    } catch (e) {
      state.session = {
        user: null,
      };
    }

    state.cookie = {
      accessToken: Cookie.accessToken(),
      refreshToken: Cookie.refreshToken(),
    };

    return state;
  });

  const updateSessionUser = (user: components["schemas"]["User"]) => {
    localStorage.setItem("user", JSON.stringify(user));
    setState({ ...state, session: { user } });
  };

  const global = {
    state,
    updateSessionUser,
  };

  return (
    <GlobalStateContext.Provider value={global}>
      {children}
    </GlobalStateContext.Provider>
  );
};
