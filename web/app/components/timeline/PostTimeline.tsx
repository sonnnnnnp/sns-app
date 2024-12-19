import { useCallback, useEffect, useState } from "react";
import client from "~/api";
import type { components } from "~/api/client";
import { Post } from "../posts/Post";

export function PostTimeline({ type }: { type: "following" | "public" }) {
  const [posts, setPosts] = useState<components["schemas"]["Post"][]>([]);

  const fetchPosts = useCallback(async (type: "following" | "public") => {
    const { data } =
      type === "following"
        ? await client.GET("/timeline/following", {})
        : await client.GET("/timeline", {});

    if (!data?.ok) {
      throw new Error("Failed to fetch post timeline");
    }

    setPosts(data.data.posts);
  }, []);

  useEffect(() => {
    fetchPosts(type);
  }, [type, fetchPosts]);

  return (
    <>
      {posts.map((post, _) => (
        <Post key={post.id} post={post} />
      ))}
    </>
  );
}
