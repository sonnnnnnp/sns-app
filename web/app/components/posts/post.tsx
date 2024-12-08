import { Avatar, AvatarIcon, Button } from "@nextui-org/react";
import {
  EllipsisVerticalIcon,
  HeartIcon,
  MessageCircleIcon,
  Repeat2Icon,
} from "lucide-react";

export function PostAction({
  icon,
  label,
  ariaLabel,
}: { icon: React.ReactNode; label?: string; ariaLabel: string }) {
  return (
    <span className="relative flex items-center text-xs text-foreground-400">
      <Button
        isIconOnly
        variant="light"
        aria-label={ariaLabel}
        className="rounded-full h-8 w-8 min-w-8"
      >
        {icon}
      </Button>
      {label && <span className="absolute left-7">{label}</span>}
    </span>
  );
}

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
      <div className="flex justify-between ml-10">
        <PostAction
          label="4"
          ariaLabel="reply"
          icon={
            <MessageCircleIcon className="h-3.5 w-3.5 text-foreground-400" />
          }
        />
        <PostAction
          ariaLabel="repost"
          icon={<Repeat2Icon className="h-3.5 w-3.5 text-foreground-400" />}
        />
        <PostAction
          label="18"
          ariaLabel="favorite"
          icon={<HeartIcon className="h-3.5 w-3.5 text-foreground-400" />}
        />
        <PostAction
          ariaLabel="options"
          icon={
            <EllipsisVerticalIcon className="h-3.5 w-3.5 text-foreground-400" />
          }
        />
      </div>
    </div>
  );
}
