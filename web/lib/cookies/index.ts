import Cookies from "js-cookie";

const Cookie = {
  accessToken(): string | undefined {
    return Cookies.get("access-token");
  },

  setAccessToken(token: string) {
    Cookies.set("access-token", token, {
      expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
    });
  },

  refreshToken(): string | undefined {
    return Cookies.get("refresh-token");
  },

  setRefreshToken(token: string) {
    Cookies.set("refresh-token", token, {
      expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
    });
  },

  destroy(): void {
    for (const cookie in Cookies.get()) {
      Cookies.remove(cookie);
    }
  },
};

export default Cookie;
