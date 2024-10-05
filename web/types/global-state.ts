import { components } from "@/lib/api/client";

export type Session = {
  user: components["schemas"]["User"] | null;
};

export type Cookie = {
  accessToken: string | null;
  refreshToken: string | null;
};

export type GlobalState = {
  session: Session | null;
  cookie: Cookie | null;
};
