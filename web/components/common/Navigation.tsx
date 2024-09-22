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
    <div className="w-52 hidden md:block">
      <div className="fixed">
        <nav className="text-xl">
          <ul>
            {navItems.map((item) => (
              <li
                key={item.label}
                className={`w-52 p-3  rounded-xl ${
                  pathname === item.href
                    ? "bg-slate-100 dark:bg-neutral-900 hover:bg-slate-200"
                    : "hover:bg-slate-100 dark:hover:bg-neutral-900"
                }`}
              >
                <Link
                  href={item.href}
                  className={`flex items-center ${
                    pathname === item.href ? "font-bold" : ""
                  }`}
                >
                  <item.icon
                    size={24}
                    strokeWidth={pathname === item.href ? "2.5" : "2"}
                  />
                  <span className="ml-4">{item.label}</span>
                </Link>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </div>
  );
}
