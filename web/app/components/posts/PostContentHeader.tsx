import { Avatar } from "@nextui-org/react";

export function PostContentHeader({ showAvatar }: { showAvatar?: boolean }) {
  return (
    <div className="flex items-center gap-1.5 mb-1">
      {showAvatar && (
        <Avatar
          src="https://i.pravatar.cc/150?u=a04258114e29026702d"
          classNames={{ base: "flex-shrink-0 w-7 h-7" }}
        />
      )}
      <span className="font-bold">ユーザー名</span>
      <span className="text-foreground-400">@username</span>
      <span className="ml-auto text-xs text-foreground-400">3時間前</span>
    </div>
  );
}
