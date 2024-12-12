import { Avatar, Badge, Tab, Tabs } from "@nextui-org/react";
import { BellIcon, Grid2x2CheckIcon, HomeIcon, MailIcon } from "lucide-react";

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
        <Tab key="/home" href="/home" title={<HomeIcon />} />
        <Tab key="/groups" href="/groups" title={<Grid2x2CheckIcon />} />
        <Tab key="/messages" href="/messages" title={<MailIcon />} />
        <Tab
          key="/notifications"
          href="/notifications"
          title={
            <Badge color="primary" size="sm" content="">
              <BellIcon />
            </Badge>
          }
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
        />
      </Tabs>
    </div>
  );
}
