import { useLocation } from "react-router";
import { LargeNavigation } from "./LargeNavigation";
import { MobileNavigation } from "./MobileNavigation";
import { SmallNavigation } from "./SmallNavigation";

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
