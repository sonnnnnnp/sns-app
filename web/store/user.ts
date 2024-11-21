import { components } from "@/lib/api/client";
import { create } from "zustand";

interface UserStore {
  user?: components["schemas"]["User"];

  setUser: (user: components["schemas"]["User"] | undefined) => void;
  clearUser: () => void;
}

export const useUserStore = create<UserStore>((set) => ({
  user: undefined,
  setUser: (user) => set({ user }),
  clearUser: () => set({ user: undefined }),
}));
