import { Outlet } from "react-router";

export default function Layout() {
  return (
    <main className="flex justify-center min-h-dvh">
      <Outlet />
    </main>
  );
}
