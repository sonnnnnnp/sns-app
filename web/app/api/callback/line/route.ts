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
        code: code!,
      },
    },
  });

  const cookieStore = cookies();
  cookieStore.set("user-id", data?.data.user_id!, {
    expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
  });
  cookieStore.set("access-token", data?.data.access_token!, {
    expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
  });
  cookieStore.set("refresh-token", data?.data.refresh_token!, {
    expires: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days
  });

  redirect("/home");
}
