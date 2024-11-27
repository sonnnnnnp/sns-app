"use client";

import * as crypto from "crypto";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Vortex } from "@/components/ui/vortex";
import client from "@/lib/api";
import { Cookie } from "@/lib/cookies";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { toast } from "sonner";

export default function Login() {
  const router = useRouter();
  const [username, setUsername] = useState("");

  const onUsernameSubmit = async () => {
    toast.promise(
      new Promise(async (resolve, reject) => {
        try {
          const { data } = await client.POST("/authorize/username", {
            params: {
              query: {
                name: username,
              },
            },
          });

          if (!data?.ok) {
            reject(`ユーザーが見つかりません。`);
          } else {
            Cookie.setAccessToken(data.data.access_token!);
            Cookie.setRefreshToken(data.data.refresh_token!);
            router.push("/home");
            resolve({ name: username });
          }
        } catch (error) {
          reject("エラーが発生しました。");
        }
      }),
      {
        loading: "ローディング中...",
        success: (data) => `ログインしました。`,
        error: (err) => `${err}`,
      }
    );
  };

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
    <div className="flex flex-col h-dvh justify-center items-center">
      <Vortex
        backgroundColor="black"
        className="flex items-center flex-col justify-center h-dvh"
      >
        <div className="grid gap-6 w-[350px]">
          <div className="flex flex-col space-y-2 text-center">
            <h1 className="text-2xl font-semibold tracking-tight">Reverie</h1>
            <p className="text-sm text-muted-foreground">
              現実の外で、もっとリアルなつながりを。
            </p>
          </div>
          {process.env.NEXT_PUBLIC_APP_ENV === "development" && (
            <div>
              <div className="grid gap-2">
                <div className="grid gap-1">
                  <Label className="sr-only" htmlFor="username">
                    Username
                  </Label>
                  <Input
                    id="username"
                    placeholder="ユーザー名"
                    type="text"
                    value={username}
                    onChange={(event) => setUsername(event.target.value)}
                  />
                </div>
                <Button onClick={onUsernameSubmit}>ログイン</Button>
              </div>
              <div className="relative mt-6">
                <div className="absolute inset-0 flex items-center">
                  <span className="w-full border-t" />
                </div>
                <div className="relative flex justify-center text-xs uppercase">
                  <span className="bg-background px-2 text-muted-foreground">
                    もしくは
                  </span>
                </div>
              </div>
            </div>
          )}

          <Button variant="outline" onClick={handleLoginWithLine}>
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img
              src="/social-media/line.svg"
              alt="LINE"
              className="mr-2 w-4 h-4"
            />
            LINEでログイン
          </Button>
        </div>
      </Vortex>
    </div>
  );
}
