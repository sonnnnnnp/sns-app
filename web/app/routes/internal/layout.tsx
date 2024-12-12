import { Outlet } from "react-router";
import { Navigation } from "~/components/shared/navigations/navigation";
import { WidgetList } from "~/components/shared/widgets/widget-list";
import { AuthProvider } from "~/providers/auth-provider";

export default function Layout() {
  return (
    <AuthProvider>
      <div className="flex justify-center min-h-dvh">
        <Navigation />
        <main className="flex-grow max-w-[750px] min-w-0 mb-16 md:mb-0">
          <Outlet />
        </main>
        <WidgetList />
      </div>
    </AuthProvider>
  );
}
