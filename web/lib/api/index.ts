import createClient, { Middleware } from "openapi-fetch";

import type { paths } from "./client";

const middleware: Middleware = {
  async onRequest({ request, options }) {},
  async onResponse({ request, response, options }) {},
};

const client = createClient<paths>({
  baseUrl: process.env.NEXT_PUBLIC_API_URL,
});

client.use(middleware);

export default client;
