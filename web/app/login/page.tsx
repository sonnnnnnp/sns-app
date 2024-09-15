"use client";

import * as crypto from "crypto";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function Login() {
  const router = useRouter();

  const handleLoginWithLine = async () => {
    const params = {
      response_type: "code",
      client_id: process.env.NEXT_PUBLIC_LINE_CHANNEL_ID,
      redirect_uri: process.env.NEXT_PUBLIC_CLIENT_URL + "/api/callback/line",
      state: crypto.randomBytes(20).toString("hex"),
      scope: "profile",
    };

    router.push(
      `https://access.line.me/oauth2/v2.1/authorize?response_type=${params.response_type}&client_id=${params.client_id}&redirect_uri=${params.redirect_uri}&state=${params.state}&scope=${params.scope}`
    );
  };

  return (
    <div className="flex h-dvh justify-center items-center">
      <Button onClick={handleLoginWithLine}>Lineでログイン</Button>
    </div>
  );
}
