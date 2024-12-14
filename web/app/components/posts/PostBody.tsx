import { Avatar, AvatarIcon } from "@nextui-org/react";
import { PostContentBody } from "./PostContentBody";
import { PostContentHeader } from "./PostContentHeader";

export function PostBody() {
  return (
    <div className="flex pb-1 text-sm">
      <Avatar icon={<AvatarIcon />} classNames={{ base: "flex-shrink-0" }} />
      <div className="grid w-full ml-2">
        <PostContentHeader />
        <PostContentBody />
      </div>
    </div>
  );
}
