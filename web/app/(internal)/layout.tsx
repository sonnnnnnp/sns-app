import { InitWrapper } from "@/components/common/init-wrapper";
import { Navigation } from "@/components/common/navigation";
import { Widget } from "@/components/common/widget";
import { Toaster } from "@/components/ui/sonner";
import { StreamProvider } from "@/providers/stream-provider";
import { Metadata } from "next";
import { ThemeProvider } from "next-themes";
import { Inter } from "next/font/google";
import React from "react";
import "../globals.css";

const inter = Inter({ subsets: ["latin"] });

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
      <body className={`${inter.className} dark:bg-black dark:text-white`}>
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
      </body>
    </html>
  );
}
