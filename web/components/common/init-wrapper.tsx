"use client";

import client from "@/lib/api";
import { useUserStore } from "@/store/user";
import { redirect } from "next/navigation";
import React from "react";

export function InitWrapper({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  React.useEffect(() => {
    (async () => {
      const { data } = await client.GET("/users/me");

      if (!data?.ok) {
        redirect("/login");
      }

      useUserStore.getState().setUser(data.data);
    })();
  }, []);

  return <>{children}</>;
}
