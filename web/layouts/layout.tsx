import HeaderNavigation from "@/components/common/header-navigation";
import FooterNavigation from "@/components/common/footer-navigation";
import { ThemeProvider } from "@/components/theme-provider";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <ThemeProvider
      attribute="class"
      defaultTheme="light"
      enableSystem
      disableTransitionOnChange
    >
      <div className="md:flex md:justify-center bg-muted/40">
        <div>
          <div className="flex min-h-screen w-full flex-col">
            <HeaderNavigation>{children}</HeaderNavigation>
            <FooterNavigation />
          </div>
        </div>
      </div>
    </ThemeProvider>
  );
}
