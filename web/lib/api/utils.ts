import client from ".";
import { Cookie } from "../cookies";

export const refreshAuthorization = async (
  refreshToken: string
): Promise<{ ok: boolean; access_token?: string; refresh_token?: string }> => {
  const { data } = await client.POST("/authorize/refresh", {
    body: {
      refresh_token: refreshToken,
    },
  });

  if (data?.ok) {
    const { access_token, refresh_token } = data.data;
    return { ok: true, access_token, refresh_token };
  }

  return { ok: false, access_token: undefined, refresh_token: undefined };
};

export const destroyAuthorization = () => {
  Cookie.destroy();

  window.location.href = "/login";
};
