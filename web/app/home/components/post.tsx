"use client";

import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import client from "@/lib/api";
import { components } from "@/lib/api/client";
import { getTimeAgo } from "@/utils/date";
import { convertNewlinesToBr } from "@/utils/text";
import { Heart, MessageCircle, Repeat2 } from "lucide-react";
import Link from "next/link";
import { useState } from "react";
import { PostHandler } from "./post-handler";

type Props = {
  post: components["schemas"]["Post"];
};

export function Post({ post }: Props) {
  const favoriteSound = new Audio("/audio/favorite.wav");

  const [favoritesCount, setFavoritesCount] = useState(post.favorites_count);
  const [isFavorited, setIsFavorited] = useState(post.favorited);

  const playFavoriteSound = () => {
    favoriteSound.volume = 0.2;
    favoriteSound.play();
  };

  const handleFavorite = async () => {
    if (isFavorited) {
      await client.DELETE("/posts/{post_id}/favorites", {
        params: {
          path: {
            post_id: post.id,
          },
        },
      });
      setFavoritesCount((prevCount) => prevCount - 1);
    } else {
      playFavoriteSound();

      await client.POST("/posts/{post_id}/favorites", {
        params: {
          path: {
            post_id: post.id,
          },
        },
      });
      setFavoritesCount((prevCount) => prevCount + 1);
    }

    setIsFavorited(!isFavorited);
  };

  return (
    <div className="flex p-2 pl-4">
      <div className="mt-2 mr-2">
        <Link href={`/users/${post.author.name}`}>
          <Avatar>
            {post.author.avatar_image_url ? (
              <AvatarImage
                src={post.author.avatar_image_url}
                className="object-cover"
              />
            ) : (
              <AvatarImage src="/users/placeholder-profile.svg" />
            )}
          </Avatar>
        </Link>
      </div>
      <div className="flex-1">
        <div className="flex items-center p-2.5">
          <div className="flex items-center gap-1.5 text-sm">
            <span className="font-semibold dark:text-white">
              {post.author.nickname}
            </span>
            <span className="text-muted-foreground dark:text-slate-400 hidden md:block">
              {`@${post.author.name}`}
            </span>
          </div>
          <span className="ml-auto text-xs text-muted-foreground dark:text-slate-400">
            {getTimeAgo(new Date(post.created_at))}
          </span>
        </div>
        <p className="px-2.5 text-sm text-muted-foreground break-all dark:text-slate-300">
          {convertNewlinesToBr(post.text ?? "")}
        </p>
        <div className="grid gap-y-1 mt-2">
          <div className="flex items-center justify-between">
            <span className="w-10">
              <Button className="p-2.5 rounded-full" variant={null} size={"sm"}>
                <MessageCircle className="h-4 w-4 text-muted-foreground opacity-70" />
                <span className="ml-1 text-muted-foreground opacity-70">
                  {/* {0} */}
                </span>
              </Button>
            </span>
            <span className="w-10">
              <Button className="p-2.5 rounded-full" variant={null} size={"sm"}>
                <Repeat2 className="h-4 w-4 text-muted-foreground opacity-70" />
                <span className="ml-1 text-muted-foreground opacity-70">
                  {/* {0} */}
                </span>
              </Button>
            </span>
            <span className="w-10">
              <Button
                className="p-2.5 rounded-full"
                variant={null}
                size={"sm"}
                onClick={handleFavorite}
              >
                <Heart className="h-4 w-4 text-muted-foreground opacity-70" />
                <span className="ml-1 text-muted-foreground opacity-70">
                  {favoritesCount ? favoritesCount : ""}
                </span>
              </Button>
            </span>
            <span className="w-10">
              <PostHandler />
              {/* <Button className="p-2.5 rounded-full" variant={null} size={"sm"}>
                <Share className="h-4 w-4 text-muted-foreground opacity-70" />
              </Button> */}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}
