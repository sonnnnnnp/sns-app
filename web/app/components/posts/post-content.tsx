import { Avatar, AvatarIcon } from "@nextui-org/react";
import { PostContentHeader } from "./post-content-header";

export function PostContent() {
  return (
    <div className="flex pb-1 text-sm">
      <Avatar icon={<AvatarIcon />} classNames={{ base: "flex-shrink-0" }} />
      <div className="grid w-full ml-2">
        <PostContentHeader />
        <span className="break-all">ここに投稿本文が入ります。</span>
      </div>
    </div>
  );
}
