"use client";

import React from "react";
import { Avatar } from "@/components/ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { Heart, MessageCircle, Repeat2, Share } from "lucide-react";
import { PostHandler } from "./post-handler";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { getTimeAgo } from "@/utils/date";
import { components } from "@/lib/api/client";
import { useState } from "react";
import client from "@/lib/api";

export const convertNewlinesToBr = (text: string) => {
  const lines = text.split("\n");
  return lines.map((line, index) => (
    <React.Fragment key={index}>
      {line}
      {index < lines.length - 1 && <br />}
    </React.Fragment>
  ));
};

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
      await client.POST("/posts/favorites/delete", {
        body: {
          post_id: post.id,
        },
      });
      setFavoritesCount((prevCount) => prevCount - 1);
    } else {
      playFavoriteSound();

      await client.POST("/posts/favorites/create", {
        body: {
          post_id: post.id,
        },
      });
      setFavoritesCount((prevCount) => prevCount + 1);
    }

    setIsFavorited(!isFavorited);
  };

  return (
    <div className="flex pl-4 pr-3 py-1.5">
      <div className="mt-2 mr-2">
        <Link href={`/users/${post.author.name}`}>
          <Avatar>
            {post.author.avatar_image_url ? (
              <AvatarImage src={post.author.avatar_image_url} />
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
              {post.author.nickname}
            </span>
            <span className="text-muted-foreground dark:text-slate-400 hidden md:block">
              {`@${post.author.name}`}
            </span>
            <span className="text-xs text-muted-foreground dark:text-slate-400">
              {` ･ ${getTimeAgo(new Date(post.created_at))}`}
            </span>
          </div>
          <div className="ml-auto">
            <PostHandler />
          </div>
        </div>
        <div className="grid gap-y-1">
          <p className="pl-2 pr-3 text-sm text-muted-foreground break-all dark:text-slate-300">
            {convertNewlinesToBr(post.text ?? "")}
          </p>
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
              <Button className="p-2.5 rounded-full" variant={null} size={"sm"}>
                <Share className="h-4 w-4 text-muted-foreground opacity-70" />
              </Button>
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}
