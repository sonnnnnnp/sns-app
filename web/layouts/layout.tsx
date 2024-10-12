import { HeaderNavigation } from "@/components/common/header-navigation";
import { Widget } from "@/components/common/widget";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex justify-center min-h-dvh bg-muted">
      <HeaderNavigation />
      {children}
      <Widget />
    </div>
  );
}
