import createClient, { Middleware } from "openapi-fetch";

import type { paths } from "./client";
import { destroyAuthorization, refreshAuthorization } from "./utils";
import Cookie from "../cookies";

interface responseBody {
  ok: boolean;
  code: number;
  data: any;
}

const UNPROTECTED_ROUTES = ["/authorize"];

const requestInterceptor: Middleware = {
  async onRequest({ request, schemaPath, options }) {
    if (
      UNPROTECTED_ROUTES.some((pathname) => schemaPath.startsWith(pathname))
    ) {
      return undefined;
    }

    request.headers.set("Authorization", `Bearer ${Cookie.accessToken()}`);
  },

  async onResponse({ request, response, options }) {
    const body: responseBody = await response.clone().json();

    if (body.code === -1002) {
      const refreshToken = Cookie.refreshToken();
      if (!refreshToken) {
        destroyAuthorization();
        return undefined;
      }

      const refreshResponse = await refreshAuthorization(refreshToken);
      if (refreshResponse.ok) {
        Cookie.setAccessToken(refreshResponse.access_token!);
        Cookie.setAccessToken(refreshResponse.refresh_token!);
        return undefined;
      }

      destroyAuthorization();
    }
  },
};

const client = createClient<paths>({
  baseUrl: process.env.NEXT_PUBLIC_API_URL,
});

client.use(requestInterceptor);

export default client;
