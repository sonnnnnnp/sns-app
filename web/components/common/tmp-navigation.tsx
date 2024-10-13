"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import {
  Home,
  Search,
  Phone,
  Video,
  MessageSquare,
  Bell,
  Settings,
} from "lucide-react";
import { Separator } from "../ui/separator";

export default function Navigation() {
  const pathname = usePathname();
  const navItems = [
    { href: "/home", label: "ホーム", icon: Home },
    { href: "/search", label: "検索", icon: Search },
    { href: "/calls", label: "通話", icon: Phone },
    { href: "/live", label: "ライブ", icon: Video },
    { href: "/messages", label: "メッセージ", icon: MessageSquare },
    { href: "/notifications", label: "通知", icon: Bell },
    { href: "/settings", label: "設定", icon: Settings },
  ];

  return (
    <div className="w-52 mr-12">
      <div className="fixed">
        <nav className="text-xl">
          <ul className="flex flex-col gap-y-5">
            {navItems.map((item) => (
              <li
                key={item.label}
                className={`w-52 rounded-xl ${
                  pathname === item.href
                    ? "bg-muted p-1 text-primary"
                    : "text-muted-foreground"
                }`}
              >
                <Link
                  href={item.href}
                  className={`flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary ${
                    pathname === item.href ? "font-[550]" : ""
                  }`}
                >
                  <item.icon
                    size={20}
                    strokeWidth={pathname === item.href ? "2.5" : "2"}
                  />
                  {item.label}
                </Link>
              </li>
            ))}
          </ul>
        </nav>
        {/* <Separator className="ml-[230px] h-[480px]" orientation="vertical" /> */}
      </div>
    </div>
  );
}
