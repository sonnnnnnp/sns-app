import { Avatar } from "@nextui-org/avatar";
import { Badge } from "@nextui-org/badge";
import { Tab, Tabs } from "@nextui-org/tabs";
import { BellIcon, Grid2x2CheckIcon, HomeIcon, MailIcon } from "lucide-react";
import ReactRouterLink from "~/utils/temp/react-router-link";

export function MobileNavigation({ pathname }: { pathname: string }) {
  return (
    <div className="fixed bottom-0 w-full z-10 backdrop-blur-md border-t md:hidden">
      <Tabs
        fullWidth
        selectedKey={pathname}
        variant="light"
        classNames={{
          tabList: "h-16",
          tab: "h-full data-[focus-visible=true]:outline-0",
          tabContent: "flex items-center",
          cursor: "hidden",
        }}
      >
        <Tab
          key="/home"
          href="/home"
          title={<HomeIcon />}
          as={ReactRouterLink}
        />
        <Tab
          key="/groups"
          href="/groups"
          title={<Grid2x2CheckIcon />}
          as={ReactRouterLink}
        />
        <Tab
          key="/messages"
          href="/messages"
          title={<MailIcon />}
          as={ReactRouterLink}
        />
        <Tab
          key="/notifications"
          href="/notifications"
          title={
            <Badge color="primary" size="sm" content="">
              <BellIcon />
            </Badge>
          }
          as={ReactRouterLink}
        />
        <Tab
          key="/users"
          href="/users"
          title={
            <Avatar
              isBordered
              size="sm"
              src="https://i.pravatar.cc/150?u=a04258114e29026702d"
            />
          }
          as={ReactRouterLink}
        />
      </Tabs>
    </div>
  );
}
