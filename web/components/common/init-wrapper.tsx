"use client";

import React from "react";

export function InitWrapper({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  React.useEffect(() => {
    console.log("InitWrapper");
  }, []);

  return <>{children}</>;
}
