"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navigation() {
  const pathname = usePathname();
  const navItems = [
    { href: "/home", label: "Home" },
    { href: "/notifications", label: "Notifications" },
    { href: "/settings", label: "Settings" },
    { href: "/live", label: "Live" },
    { href: "/messages", label: "Messages" },
    { href: "/calls", label: "Calls" },
  ];

  return (
    <div className="w-52">
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
