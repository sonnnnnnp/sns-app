import createClient, { Middleware } from "openapi-fetch";

import { Cookie } from "../cookies";
import type { paths } from "./client";
import { destroyAuthorization, refreshAuthorization } from "./utils";

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

    if (body.code === -110002) {
      const refreshToken = Cookie.refreshToken();
      if (!refreshToken) {
        destroyAuthorization();
        return undefined;
      }

      const refreshResponse = await refreshAuthorization(refreshToken);
      if (!refreshResponse.ok) {
        destroyAuthorization();
        return undefined;
      }

      Cookie.setAccessToken(refreshResponse.access_token!);
      Cookie.setRefreshToken(refreshResponse.refresh_token!);

      // retry original request

      request.headers.set(
        "Authorization",
        `Bearer ${refreshResponse.access_token!}`
      );

      return fetch(request);
    }
  },
};

const client = createClient<paths>({
  baseUrl: process.env.NEXT_PUBLIC_API_URL,
});

client.use(requestInterceptor);

export default client;
