import { Navigation } from "@/components/common/navigation";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex justify-center min-h-dvh bg-muted">
      <Navigation />
      {children}
    </div>
  );
}
