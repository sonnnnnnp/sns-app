import { useLocation } from "react-router";
import { LargeNavigation } from "./large-navigation";
import { MobileNavigation } from "./mobile-navigation";
import { SmallNavigation } from "./small-navigation";

export function Navigation() {
  const { pathname } = useLocation();

  return (
    <div>
      <LargeNavigation pathname={pathname} />
      <SmallNavigation pathname={pathname} />
      <MobileNavigation pathname={pathname} />
    </div>
  );
}
