import { Avatar, Badge, Listbox, ListboxItem } from "@nextui-org/react";
import {
  BellIcon,
  Grid2x2CheckIcon,
  HomeIcon,
  MailIcon,
  SearchIcon,
  SettingsIcon,
} from "lucide-react";
import ReactRouterLink from "~/utils/temp/react-router-link";

export function SmallNavigation({ pathname }: { pathname: string }) {
  return (
    <div className="sticky inset-y-0 h-dvh w-16 px-2 py-4 flex-shrink-0 border-r hidden md:block lg:hidden">
      <Listbox
        aria-label="Actions"
        selectionMode="single"
        selectedKeys={[pathname]}
        className="h-full [&>ul]:h-full"
        itemClasses={{
          base: "w-10 h-10 rounded-full transition-colors mb-4 text-default-500 data-[selected=true]:bg-default data-[selected=true]:text-default-foreground last:mb-0",
          title: "hidden",
          selectedIcon: "hidden",
        }}
      >
        <ListboxItem
          key="/home"
          href="/home"
          title="ホーム"
          startContent={<HomeIcon />}
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/explore"
          href="/explore"
          title="検索"
          startContent={<SearchIcon />}
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/groups"
          href="/groups"
          title="グループ"
          startContent={<Grid2x2CheckIcon />}
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/messages"
          href="/messages"
          title="メッセージ"
          startContent={<MailIcon />}
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/notifications"
          href="/notifications"
          title="通知"
          startContent={
            <Badge color="primary" size="sm" content="">
              <BellIcon />
            </Badge>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/settings"
          href="/settings"
          title="設定"
          startContent={<SettingsIcon />}
          className="mt-auto"
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/users"
          href="/users"
          title="プロフィール"
          startContent={
            <Avatar
              isBordered
              src="https://i.pravatar.cc/150?u=a04258114e29026702d"
              className="w-6 h-6"
            />
          }
          as={ReactRouterLink}
        />
      </Listbox>
    </div>
  );
}
