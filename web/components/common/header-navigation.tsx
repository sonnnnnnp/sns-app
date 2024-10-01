"use client";

import Image from "next/image";
import Link from "next/link";
import {
  Bell,
  Home,
  LoaderPinwheel,
  MessagesSquare,
  PanelLeft,
  Phone,
  Search,
  Settings,
} from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Sheet,
  SheetContent,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { useContext, useEffect } from "react";
import client from "@/lib/api";
import { GlobalStateContext } from "@/context/global-state-provider";

export const iframeHeight = "938px";

export const containerClassName = "w-full h-full";

type Props = {
  children: React.ReactNode;
};

export default function HeaderNavigation({ children }: Props) {
  const context = useContext(GlobalStateContext);

  useEffect(() => {
    const fetchMyself = async () => {
      const { data } = await client.GET("/users/me");
      if (!data?.ok) {
        return console.error("error fetching myself");
      }
      context.updateSessionUser(data?.data);
    };

    console.log(JSON.stringify(context.state));

    if (!context.state.session?.user) {
      fetchMyself();
    }
  }, [context]);

  return (
    <div>
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
                  href="/calls"
                  className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                >
                  <Phone className="h-5 w-5" />
                  <span className="sr-only">通話</span>
                </Link>
              </TooltipTrigger>
              <TooltipContent side="right">通話</TooltipContent>
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
        </nav>
      </aside>
      <div className="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
        <header className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6">
          <Sheet>
            <SheetTrigger asChild>
              <Button size="icon" variant="outline" className="sm:hidden">
                <PanelLeft className="h-5 w-5" />
                <span className="sr-only">Toggle Menu</span>
              </Button>
            </SheetTrigger>
            <SheetContent side="left" className="sm:max-w-xs">
              <nav className="grid gap-6 text-lg font-medium">
                <SheetTitle>
                  <Link
                    href="/home"
                    className="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:text-base"
                  >
                    <LoaderPinwheel className="h-5 w-5 transition-all group-hover:scale-110" />
                    <span className="sr-only">sns-app</span>
                  </Link>
                </SheetTitle>
                <Link
                  href="/home"
                  className="flex items-center gap-4 px-2.5 text-foreground"
                >
                  <Home className="h-5 w-5" />
                  ホーム
                </Link>
                <Link
                  href="/calls"
                  className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                >
                  <Phone className="h-5 w-5" />
                  通話
                </Link>
                <Link
                  href="/messages"
                  className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                >
                  <MessagesSquare className="h-5 w-5" />
                  メッセージ
                </Link>
                <Link
                  href="/notifications"
                  className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                >
                  <Bell className="h-5 w-5" />
                  通知
                </Link>
                <Link
                  href="/settings"
                  className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                >
                  <Settings className="h-5 w-5" />
                  設定
                </Link>
              </nav>
            </SheetContent>
          </Sheet>
          <div className="relative mr-auto flex-1 md:grow-0">
            <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              type="search"
              placeholder="Search..."
              className="w-full rounded-lg bg-background pl-8 md:w-[200px] lg:w-[336px]"
            />
          </div>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                size="icon"
                className="overflow-hidden rounded-full"
              >
                {context.state.session?.user?.avatar_url ? (
                  <Image
                    src={context.state.session.user.avatar_url}
                    width={36}
                    height={36}
                    alt="Avatar"
                    className="overflow-hidden rounded-full"
                  />
                ) : (
                  <Image
                    src="/users/placeholder-profile.svg"
                    width={36}
                    height={36}
                    alt="Avatar"
                    className="overflow-hidden rounded-full"
                  />
                )}
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <Link href={`/users/${context.state.session?.user?.id}`}>
                <DropdownMenuLabel>
                  {context.state.session?.user?.display_name === ""
                    ? "未設定"
                    : context.state.session?.user?.display_name}
                </DropdownMenuLabel>
              </Link>
              <DropdownMenuSeparator />
              <DropdownMenuItem>通知</DropdownMenuItem>
              <DropdownMenuItem>設定</DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem>ログアウト</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </header>
        <main className="p-4 sm:px-6 sm:py-0">{children}</main>
      </div>
    </div>
  );
}
