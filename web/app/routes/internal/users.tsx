import { Avatar, Button, Tab, Tabs } from "@nextui-org/react";
import {
  ArrowLeftIcon,
  EllipsisVerticalIcon,
  MailIcon,
  SearchIcon,
} from "lucide-react";
import { Link, useNavigate } from "react-router";
import { Post } from "~/components/posts/post";

export default function User() {
  const navigate = useNavigate();

  return (
    <div className="w-full">
      <div className="sticky top-0 z-10 flex items-center px-2 w-full h-14 bg-transparent backdrop-blur-md border-b">
        <Button
          isIconOnly
          variant="light"
          radius="full"
          onPress={() => navigate(-1)}
        >
          <ArrowLeftIcon className="w-5 h-5" />
        </Button>
        <div className="grid gap-1 ml-2">
          <span className="font-bold">ユーザー名</span>
          <span className="text-xs text-foreground-400">214件の投稿</span>
        </div>
        <Button isIconOnly variant="light" radius="full" className="ml-auto">
          <EllipsisVerticalIcon className="w-5 h-5" />
        </Button>
      </div>
      <div className="relative w-full">
        <img
          src="https://nextui.org/images/hero-card-complete.jpeg"
          alt="banner-image"
          className="object-cover w-full h-40"
        />
        <Avatar
          className="absolute -bottom-10 left-4 w-24 h-24"
          src="https://i.pravatar.cc/150?u=a04258114e29026708c"
        />
      </div>
      <div className="flex justify-end space-x-2 pt-4 px-6">
        <Button
          isIconOnly
          variant="bordered"
          radius="full"
          className="h-8 w-8 min-w-8"
        >
          <SearchIcon className="w-4 h-4" />
        </Button>
        <Button
          isIconOnly
          variant="bordered"
          radius="full"
          className="h-8 w-8 min-w-8"
        >
          <MailIcon className="w-4 h-4" />
        </Button>
        <Button radius="full" className="h-8 bg-foreground text-default">
          フォローする
        </Button>
      </div>
      <div className="grid gap-4 px-6 pb-2 text-sm">
        <div className="flex flex-col">
          <span className="font-bold text-medium">ユーザー名</span>
          <span className="text-foreground-400">@username</span>
        </div>
        <p>
          自己紹介文がここに入ります。
          <br />
          そうです。ここです。
        </p>
        <div className="flex gap-3 text-xs">
          <Link to="/users/user_id/following" className="hover:underline">
            <span className="font-bold mr-1">24</span>
            <span className="text-foreground-400">フォロー中</span>
          </Link>
          <Link to="/users/user_id/following" className="hover:underline">
            <span className="font-bold mr-1">132</span>
            <span className="text-foreground-400">フォロワー</span>
          </Link>
        </div>
      </div>
      <Tabs
        fullWidth
        variant="underlined"
        classNames={{
          base: "sticky top-[55px] bg-background border-b z-10",
          tabList: "p-0",
          tab: "h-12",
          cursor: "w-[65%]",
          panel: "p-0",
        }}
      >
        <Tab key="posts" title="投稿">
          {Array.from({ length: 20 }).map((_, i) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
            <Post key={i} />
          ))}
        </Tab>
        <Tab key="replies" title="返信">
          {Array.from({ length: 20 }).map((_, i) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
            <Post key={i} />
          ))}
        </Tab>
        <Tab key="calls" title="通話">
          <div>通話</div>
        </Tab>
        <Tab key="media" title="メディア">
          {Array.from({ length: 20 }).map((_, i) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
            <Post key={i} />
          ))}
        </Tab>
        <Tab key="favorites" title="いいね">
          {Array.from({ length: 20 }).map((_, i) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
            <Post key={i} />
          ))}
        </Tab>
      </Tabs>
    </div>
  );
}
