import { NextRequest, NextResponse } from "next/server";

export async function middleware(req: NextRequest) {
  const path = new URL(req.url).pathname;

  // 認証が不必要なパスを除外
  if (path.startsWith("/login") || path.startsWith("/api")) {
    return NextResponse.next();
  }

  // クッキーから取得したトークンの有効性を確認
  let accessToken = req.cookies.get("access-token")?.value;
  let refreshToken = req.cookies.get("refresh-token")?.value;

  if (accessToken && refreshToken) {
    if (path === "/" || path.startsWith("/login")) {
      return NextResponse.redirect(new URL("/home", req.url));
    }
    return NextResponse.next();
  }

  return NextResponse.redirect(new URL("/login", req.url));
}

export const config = {
  matcher: [
    "/((?!_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)",
  ],
};
