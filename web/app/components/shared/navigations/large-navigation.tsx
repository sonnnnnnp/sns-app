import {
  Badge,
  Button,
  Divider,
  Listbox,
  ListboxItem,
  User,
  useDisclosure,
} from "@nextui-org/react";
import {
  BellIcon,
  Grid2x2CheckIcon,
  HomeIcon,
  MailIcon,
  PencilIcon,
  SearchIcon,
  SettingsIcon,
  UserRoundIcon,
} from "lucide-react";
import { PostModal } from "~/components/posts/post-modal";
import ReactRouterLink from "~/utils/temp/react-router-link";

export function LargeNavigation({ pathname }: { pathname: string }) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <div className="sticky inset-y-0 h-dvh w-24 p-2 flex-shrink-0 border-r hidden lg:block lg:w-64 lg:p-6">
      <div className="flex flex-col gap-5 px-2 pt-4">
        <User
          className="justify-start"
          classNames={{ name: "font-bold" }}
          avatarProps={{
            src: "https://i.pravatar.cc/150?u=a04258114e29026702d",
          }}
          description="@username"
          name="ユーザー名"
        />
        <Button
          onPress={onOpen}
          className="rounded-full bg-foreground text-default"
          endContent={<PencilIcon className="w-4 h-4" />}
        >
          投稿する
        </Button>
        <PostModal isOpen={isOpen} onOpenChange={onOpenChange} />
      </div>
      <Divider className="my-4" />
      <Listbox
        aria-label="Actions"
        selectionMode="single"
        selectedKeys={[pathname]}
        itemClasses={{
          base: "w-full h-12 rounded-full transition-colors mb-2 text-default-500 data-[selected=true]:bg-default data-[selected=true]:text-default-foreground last:mb-0",
          title: "hidden",
          selectedIcon: "hidden",
        }}
      >
        <ListboxItem
          key="/home"
          href="/home"
          title="ホーム"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <HomeIcon />
              ホーム
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/explore"
          href="/explore"
          title="見つける"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <SearchIcon />
              見つける
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/groups"
          href="/groups"
          title="グループ"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <Grid2x2CheckIcon />
              グループ
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/messages"
          href="/messages"
          title="メッセージ"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <MailIcon />
              メッセージ
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/notifications"
          href="/notifications"
          title="通知"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <Badge color="primary" size="sm" content="">
                <BellIcon />
              </Badge>
              通知
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/users"
          href="/users"
          title="プロフィール"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <UserRoundIcon />
              プロフィール
            </span>
          }
          as={ReactRouterLink}
        />
        <ListboxItem
          key="/settings"
          href="/settings"
          title="設定"
          startContent={
            <span className="flex items-center gap-4 ml-2 text-small font-medium">
              <SettingsIcon />
              設定
            </span>
          }
          as={ReactRouterLink}
        />
      </Listbox>
    </div>
  );
}
