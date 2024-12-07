import { Avatar, AvatarIcon, Button } from "@nextui-org/react";
import {
  EllipsisVerticalIcon,
  HeartIcon,
  MessageCircleIcon,
  Repeat2Icon,
} from "lucide-react";

export function Post() {
  return (
    <div className="flex flex-col p-4 pb-1 border-b last:border-none">
      <div className="flex pb-1 text-sm">
        <Avatar icon={<AvatarIcon />} classNames={{ base: "flex-shrink-0" }} />
        <div className="grid w-full ml-2">
          <div className="flex gap-1.5 mb-1">
            <span className="font-bold">ユーザー名</span>
            <span className="text-foreground-400">@username</span>
            <span className="ml-auto text-xs text-foreground-400">3時間前</span>
          </div>
          <span className="break-all">ここに投稿本文が入ります。</span>
        </div>
      </div>
      <div className="flex justify-between ml-10 text-xs text-foreground-400">
        <span className="relative flex items-center">
          <Button
            isIconOnly
            variant="light"
            aria-label="reply"
            className="rounded-full h-8 w-8 min-w-8"
          >
            <MessageCircleIcon className="h-3.5 w-3.5 text-foreground-400" />
          </Button>
          <span className="absolute left-7">4</span>
        </span>
        <span className="relative flex items-center">
          <Button
            isIconOnly
            variant="light"
            aria-label="repost"
            className="rounded-full h-8 w-8 min-w-8"
          >
            <Repeat2Icon className="h-3.5 w-3.5 text-foreground-400" />
          </Button>
        </span>
        <span className="relative flex items-center">
          <Button
            isIconOnly
            variant="light"
            aria-label="favorite"
            className="rounded-full h-8 w-8 min-w-8"
          >
            <HeartIcon className="h-3.5 w-3.5 text-foreground-400" />
          </Button>
          <span className="absolute left-7">34</span>
        </span>
        <span className="relative flex items-center">
          <Button
            isIconOnly
            variant="light"
            aria-label="options"
            className="rounded-full h-8 w-8 min-w-8"
          >
            <EllipsisVerticalIcon className="h-3.5 w-3.5 text-foreground-400" />
          </Button>
        </span>
      </div>
    </div>
  );
}
