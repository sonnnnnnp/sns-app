import { Avatar } from "@/components/ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { Heart, MessageCircle, Repeat2, Share } from "lucide-react";
import { PostHandler } from "./post-handler";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { getTimeAgo } from "@/utils/date";

export function Post(props: {
  username: string;
  display_name: string;
  avatar_image_url: string;
  content: string;
  created_at: string;
}) {
  return (
    <div className="flex border-b pl-4 pr-3 py-1.5">
      <div className="mt-2 mr-2">
        <Link href={`/users/${props.username}`}>
          <Avatar>
            {props.avatar_image_url ? (
              <AvatarImage src={props.avatar_image_url} />
            ) : (
              <AvatarImage src="/users/placeholder-profile.svg" />
            )}
          </Avatar>
        </Link>
      </div>
      <div className="flex-1">
        <div className="flex items-center">
          <div className="flex items-center gap-1.5 px-2 text-sm">
            <span className="font-semibold dark:text-white">
              {props.display_name}
            </span>
            <span className="text-muted-foreground dark:text-slate-400 hidden md:block">
              {`@${props.username}`}
            </span>
            <span className="text-xs text-muted-foreground dark:text-slate-400">
              {` ･ ${getTimeAgo(new Date(props.created_at))}`}
            </span>
          </div>
          <div className="ml-auto">
            <PostHandler />
          </div>
        </div>
        <div className="grid gap-y-1">
          <p className="text-sm text-muted-foreground dark:text-slate-300 px-2">
            {props.content}
          </p>
          <div className="flex items-center justify-between">
            <Button
              className="p-2.5 rounded-full"
              variant={"ghost"}
              size={"sm"}
            >
              <MessageCircle className="className=h-4 w-4 text-muted-foreground opacity-70" />
            </Button>
            <Button
              className="p-2.5 rounded-full"
              variant={"ghost"}
              size={"sm"}
            >
              <Repeat2 className="className=h-4 w-4 text-muted-foreground opacity-70" />
            </Button>
            <Button
              className="p-2.5 rounded-full"
              variant={"ghost"}
              size={"sm"}
            >
              <Heart className="className=h-4 w-4 text-muted-foreground opacity-70" />
            </Button>
            <Button
              className="p-2.5 rounded-full"
              variant={"ghost"}
              size={"sm"}
            >
              <Share className="className=h-4 w-4 text-muted-foreground opacity-70" />
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
