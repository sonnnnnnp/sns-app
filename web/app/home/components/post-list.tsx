"use client";

import { useEffect, useState } from "react";
import { Post } from "./post";
import client from "@/lib/api";
import { components } from "@/lib/api/client";

export function PostList() {
  const [posts, setPosts] = useState<
    components["schemas"]["Timeline"]["posts"]
  >([]);

  useEffect(() => {
    const fetchTimeline = async () => {
      const { data } = await client.GET("/timeline");
      if (!data?.ok) {
        return console.error("error fetching timeline");
      }

      setPosts(data?.data.posts ?? []);
    };

    fetchTimeline();
  }, []);

  return (
    <div className="">
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
    </div>
  );
}
