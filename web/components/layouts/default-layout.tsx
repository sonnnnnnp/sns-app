import { GlobalStateProvider } from "@/context/global-state-provider";
import FooterNavigation from "../common/footer-navigation";
import HeaderNavigation from "../common/header-navigation";

export default function DefaultLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <GlobalStateProvider>
      <div className="flex min-h-screen w-full flex-col bg-muted/40">
        <HeaderNavigation>{children}</HeaderNavigation>
        <FooterNavigation />
      </div>
    </GlobalStateProvider>
  );
}
