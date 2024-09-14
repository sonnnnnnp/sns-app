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
      <CardContent className="w-full grid gap-6 pt-5 pb-2">
        <div className="space-x-3 flex">
          <div>
            <Link href={`/users/${props.username}`}>
              <Avatar>
                <AvatarImage src={props.avatar_image_url} />
              </Avatar>
            </Link>
          </div>
          <div>
            <div className="flex items-center">
              <div className="flex items-center gap-1">
                <span className="text-sm font-semibold leading-none">
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
            <div className="mt-1 mb-2">
              <p className="text-xs text-muted-foreground">{props.content}</p>
            </div>
            <div className="flex items-center justify-between opacity-60">
              <Button variant={"ghost"} size={"sm"}>
                <MessageCircle
                  className="cursor-pointer"
                  size={16}
                  strokeWidth={1.6}
                />
              </Button>
              <Button variant={"ghost"} size={"sm"}>
                <Repeat2
                  className="cursor-pointer"
                  size={16}
                  strokeWidth={1.6}
                />
              </Button>
              <Button variant={"ghost"} size={"sm"}>
                <Heart className="cursor-pointer" size={16} strokeWidth={1.6} />
              </Button>
              <Button variant={"ghost"} size={"sm"}>
                <Share className="cursor-pointer" size={16} strokeWidth={1.6} />
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
