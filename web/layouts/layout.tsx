import HeaderNavigation from "@/components/common/header-navigation";
import FooterNavigation from "@/components/common/footer-navigation";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex min-h-screen w-full flex-col bg-muted/40">
      <HeaderNavigation>{children}</HeaderNavigation>
      <FooterNavigation />
    </div>
  );
}
