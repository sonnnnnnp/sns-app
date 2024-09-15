import { client } from "@/lib/api";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { NextRequest } from "next/server";

export async function GET(req: NextRequest) {
  const searchParams = req.nextUrl.searchParams;
  const code = searchParams.get("code");

  const { data } = await client.POST("/authorize/line", {
    params: {
      query: {
        code: code as string,
      },
    },
  });

  const cookieStore = cookies();
  cookieStore.set("user-id", data?.data.user_id as string, {
    expires: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000), // 7 days
  });
  cookieStore.set("access-token", data?.data.access_token as string, {
    expires: new Date(Date.now() + 24 * 60 * 60 * 1000), // 1 day
  });
  cookieStore.set("refresh-token", data?.data.refresh_token as string, {
    expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
  });

  redirect("/home");
}
