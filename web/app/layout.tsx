import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Toaster } from "@/components/ui/sonner";
import { ThemeProvider } from "next-themes";
import { StreamProvider } from "@/providers/stream-provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "sns-app",
  description: "sns-app",
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
        className={`${inter.className} fixed top-0 left-0 w-full h-full overflow-y-scroll dark:bg-black dark:text-white`}
      >
        <div className="absolute top-0 left-0 w-full h-full overflow-scroll">
          <ThemeProvider
            attribute="class"
            defaultTheme="light"
            enableSystem
            disableTransitionOnChange
          >
            <StreamProvider>
              {children}
              <Toaster />
            </StreamProvider>
          </ThemeProvider>
        </div>
      </body>
    </html>
  );
}
