import { NextRequest, NextResponse } from "next/server";
import client from "./lib/api";

async function refreshTokens(
  refreshToken: string
): Promise<{ ok: boolean; access_token?: string; refresh_token?: string }> {
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
}

export async function middleware(req: NextRequest) {
  const path = new URL(req.url).pathname;

  // 認証が不必要なパスを除外
  if (path.startsWith("/login") || path.startsWith("/api")) {
    return NextResponse.next();
  }

  // クッキーから取得したトークンの有効性を確認
  let accessToken = req.cookies.get("access-token")?.value;
  let refreshToken = req.cookies.get("refresh-token")?.value;

  if (!accessToken) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  const { data } = await client.GET("/users/me", {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });

  console.log("Response: " + JSON.stringify(data));

  if (data?.ok) {
    return NextResponse.next();
  }

  // トークン期限切れの場合、リフレッシュトークンを使って再取得
  if (data?.code === -1002) {
    const { ok, access_token, refresh_token } = await refreshTokens(
      refreshToken!
    );

    console.log(
      "Refresh Response: " + JSON.stringify({ ok, access_token, refresh_token })
    );

    if (ok) {
      const response = NextResponse.next();

      response.cookies.set("access-token", access_token!, {
        expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
      });
      response.cookies.set("refresh-token", refresh_token!, {
        expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
      });

      return response;
    }
  }

  // 再取得に失敗した場合はクッキーを削除してログイン画面にリダイレクト
  const response = NextResponse.redirect(new URL("/login", req.url));

  req.cookies.getAll().forEach((cookie) => {
    response.cookies.delete(cookie.name);
  });

  return response;
}

export const config = {
  matcher: [
    "/((?!_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)",
  ],
};
