import { Navigation } from "@/components/common/navigation";
import { Widget } from "@/components/common/widget";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex justify-center min-h-dvh">
      <Navigation />
      <main className="flex-grow max-w-[750px]">{children}</main>
      <Widget />
    </div>
  );
}
