import { BackgroundTexture } from "@/components/common/background-texture";
import { InitWrapper } from "@/components/common/init-wrapper";
import { Navigation } from "@/components/common/navigation";
import { Widget } from "@/components/common/widget";
import { Toaster } from "@/components/ui/sonner";
import { StreamProvider } from "@/providers/stream-provider";
import { Metadata } from "next";
import { ThemeProvider } from "next-themes";
import { M_PLUS_Rounded_1c } from "next/font/google";
import React from "react";
import "../globals.css";

const mPlusRounded1c = M_PLUS_Rounded_1c({ weight: "400", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Reverie",
  description: "Social media app.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja" suppressHydrationWarning>
      <head></head>
      <body
        className={`${mPlusRounded1c.className} dark:bg-black dark:text-white`}
      >
        <div className="fixed top-0 left-0 w-full h-full overflow-y-scroll">
          <ThemeProvider
            attribute="class"
            defaultTheme="light"
            enableSystem
            disableTransitionOnChange
          >
            <StreamProvider>
              <InitWrapper>
                <div className="flex justify-center min-h-dvh">
                  <Navigation />
                  <main className="flex-grow max-w-[750px]">{children}</main>
                  <Widget />
                </div>
              </InitWrapper>
              <Toaster />
            </StreamProvider>
          </ThemeProvider>
        </div>
        <BackgroundTexture />
      </body>
    </html>
  );
}
