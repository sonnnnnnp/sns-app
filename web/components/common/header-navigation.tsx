"use client";

import Link from "next/link";
import {
  Bell,
  Home,
  LoaderPinwheel,
  MessagesSquare,
  Search,
  Settings,
  Users,
} from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Avatar, AvatarImage } from "../ui/avatar";
import { Badge } from "../ui/badge";

export const iframeHeight = "938px";

export const containerClassName = "w-full h-full";

type Props = {
  children: React.ReactNode;
};

export default function HeaderNavigation({ children }: Props) {
  return (
    <div>
      <div className="fixed inset-y-0 z-10 w-64 flex-col hidden md:block">
        <div className="flex h-full max-h-screen flex-col gap-2">
          <div className="flex h-12 items-center px-6">
            <Link href="/" className="flex items-center gap-2 font-semibold">
              <LoaderPinwheel className="h-6 w-6" />
              <span className="">sns-app</span>
            </Link>
            <Button variant="outline" size="icon" className="ml-auto h-8 w-8">
              <Bell className="h-4 w-4" />
              <span className="sr-only">Toggle notifications</span>
            </Button>
          </div>
          <div className="flex-1">
            <nav className="grid gap-y-3 items-start mt-3 px-4 text-sm font-medium">
              <Link
                href="/home"
                className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
              >
                <Home className="h-5 w-5" />
                ホーム
              </Link>
              <Link
                href="/search"
                className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
              >
                <Search className="h-5 w-5" />
                検索
              </Link>
              <Link
                href="/messages"
                className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
              >
                <MessagesSquare className="h-5 w-5" />
                メッセージ
                <Badge className="ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full">
                  6
                </Badge>
              </Link>
              <Link
                href="/groups"
                className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
              >
                <Users className="h-5 w-5" />
                グループ
              </Link>
              <Link
                href="/settings"
                className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
              >
                <Settings className="h-5 w-5" />
                設定
              </Link>
            </nav>
          </div>
          {/* <div className="mt-auto p-4">
            <Card x-chunk="dashboard-02-chunk-0">
              <CardHeader className="p-2 pt-0 md:p-4">
                <CardTitle>Ver. 0.10</CardTitle>
                <CardDescription>アップデート情報</CardDescription>
              </CardHeader>
              <CardContent className="p-2 pt-0 md:p-4 md:pt-0">
                <Button size="sm" className="w-full">
                  もっとみる
                </Button>
              </CardContent>
            </Card>
          </div> */}
        </div>
      </div>
      <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex md:hidden">
        <nav className="flex flex-col items-center gap-4 px-2 sm:py-5">
          <Link
            href="/home"
            className="group flex h-9 w-9 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:h-8 md:w-8 md:text-base"
          >
            <LoaderPinwheel className="h-4 w-4 transition-all group-hover:scale-110" />
            <span className="sr-only">sns-app</span>
          </Link>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/home"
                  className="flex h-9 w-9 items-center justify-center rounded-lg bg-accent text-accent-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Home className="h-5 w-5" />
                  <span className="sr-only">ホーム</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">ホーム</TooltipContent>
            </Tooltip>
          </TooltipProvider>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/search"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Search className="h-5 w-5" />
                  <span className="sr-only">検索</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">検索</TooltipContent>
            </Tooltip>
          </TooltipProvider>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/messages"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <MessagesSquare className="h-5 w-5" />
                  <span className="sr-only">メッセージ</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">メッセージ</TooltipContent>
            </Tooltip>
          </TooltipProvider>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/groups"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Users className="h-5 w-5" />
                  <span className="sr-only">グループ</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">グループ</TooltipContent>
            </Tooltip>
          </TooltipProvider>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/notifications"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Bell className="h-5 w-5" />
                  <span className="sr-only">通知</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">通知</TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </nav>
        <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/settings"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Settings className="h-5 w-5" />
                  <span className="sr-only">設定</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">設定</TooltipContent>
            </Tooltip>
          </TooltipProvider>
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <Link
                  href="/users/username"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Avatar className="h-8 w-8">
                    <AvatarImage src="/users/placeholder-profile.svg" />
                  </Avatar>
                  <span className="sr-only">ユーザー名</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">ユーザー名</TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </nav>
      </aside>
      <div className="flex flex-col sm:gap-4 sm:pl-14 md:pl-64">
        <main>{children}</main>
      </div>
    </div>
  );
}
