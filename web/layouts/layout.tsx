import { Navigation } from "@/components/common/navigation";
import { Widget } from "@/components/common/widget";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex min-h-dvh">
      <Navigation />
      <main className="flex-grow">{children}</main>
      <Widget />
    </div>
  );
}
