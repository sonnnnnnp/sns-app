"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navigation() {
  const pathname = usePathname();
  const navItems = [
    { href: "/home", label: "ホーム" },
    { href: "/search", label: "検索" },
    { href: "/calls", label: "通話" },
    { href: "/live", label: "ライブ" },
    { href: "/messages", label: "メッセージ" },
    { href: "/notifications", label: "通知" },
    { href: "/settings", label: "設定" },
  ];
  
  return (
    <div className="w-52 hidden md:block">
      <div className="fixed">
        <nav className="text-xl">
          <ul>
            {navItems.map((item) => (
              <li
                key={item.label}
                className="w-52 p-3 hover:bg-slate-100 dark:hover:bg-neutral-900 rounded-xl"
              >
                <Link
                  href={item.href}
                  className={`block ${
                    pathname === item.href ? "font-bold" : ""
                  }`}
                >
                  {item.label}
                </Link>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </div>
  );
}
