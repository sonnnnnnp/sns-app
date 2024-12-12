import { Avatar, AvatarIcon } from "@nextui-org/react";

export function PostContent() {
  return (
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
  );
}
