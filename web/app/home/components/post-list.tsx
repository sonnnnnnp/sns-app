"use client";

import { Post } from "./post";
import { components } from "@/lib/api/client";
import { Card, CardContent } from "@/components/ui/card";

type Props = {
  posts: components["schemas"]["Timeline"]["posts"];
};

export function PostList({ posts }: Props) {
  return (
    <Card className="flex text-sm my-3 w-full dark:bg-black dark:border-slate-800 mb-16 sm:mb-0">
      <CardContent className="w-full p-0">
        {posts.map((post, i) => (
          <Post
            key={i}
            username={post.author.username}
            display_name={post.author.display_name}
            avatar_image_url={post.author.avatar_url}
            content={post.content ?? ""}
            created_at={post.created_at}
          />
        ))}
      </CardContent>
    </Card>
  );
}
