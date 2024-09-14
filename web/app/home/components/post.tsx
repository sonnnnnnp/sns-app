import { Avatar } from "@/components/ui/avatar";
import { Card, CardContent } from "@/components/ui/card";
import { AvatarImage } from "@radix-ui/react-avatar";
import { Heart, MessageCircle, Repeat2, Share } from "lucide-react";
import { PostHandler } from "./post-handler";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export function Post(props: {
  username: string;
  display_name: string;
  avatar_image_url: string;
  content: string;
  created_at: number;
}) {
  return (
    <Card className="flex flex-col items-start gap-2 text-left text-sm md:w-[580px]">
      <CardContent className="w-full grid gap-6 p-3 py-1">
        <div className="flex">
          <div className="mt-2 mr-2">
            <Link href={`/users/${props.username}`}>
              <Avatar>
                <AvatarImage src={props.avatar_image_url} />
              </Avatar>
            </Link>
          </div>
          <div className="flex-1">
            <div className="flex items-center">
              <div className="flex items-center gap-1">
                <span className="text-sm font-semibold">
                  {props.display_name}
                </span>
                <span className="text-sm text-muted-foreground">
                  {`@${props.username}`}
                </span>
                <span className="text-xs text-muted-foreground">{` ･ 3分前`}</span>
              </div>
              <div className="ml-auto">
                <PostHandler />
              </div>
            </div>
            <div className="mb-1.5">
              <p className="text-sm text-muted-foreground">{props.content}</p>
            </div>
            <div className="flex items-center justify-between opacity-60">
              <Button
                className="p-2.5 rounded-full"
                variant={"ghost"}
                size={"sm"}
              >
                <MessageCircle
                  className="cursor-pointer"
                  size={16}
                  strokeWidth={1.6}
                />
              </Button>
              <Button
                className="p-2.5 rounded-full"
                variant={"ghost"}
                size={"sm"}
              >
                <Repeat2
                  className="cursor-pointer"
                  size={16}
                  strokeWidth={1.6}
                />
              </Button>
              <Button
                className="p-2.5 rounded-full"
                variant={"ghost"}
                size={"sm"}
              >
                <Heart className="cursor-pointer" size={16} strokeWidth={1.6} />
              </Button>
              <Button
                className="p-2.5 rounded-full"
                variant={"ghost"}
                size={"sm"}
              >
                <Share className="cursor-pointer" size={16} strokeWidth={1.6} />
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
