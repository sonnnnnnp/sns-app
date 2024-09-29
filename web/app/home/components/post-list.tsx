"use client";

import { useEffect, useState } from "react";
import { Post } from "./post";
import client from "@/lib/api";
import { components } from "@/lib/api/client";

export function PostList() {
  const [timeline, setTimeline] = useState<
    components["schemas"]["Timeline"]["posts"]
  >([]);

  useEffect(() => {
    const fetchTimeline = async () => {
      const { data } = await client.GET("/timeline");
      if (!data?.ok) {
        alert(data?.data);
      }

      setTimeline(data?.data.posts!);
    };

    fetchTimeline();
  }, []);

  return (
    <div className="">
      {[...Array(30)].map((_, i) => (
        <Post
          key={i}
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="s"
          created_at={1234567}
        />
      ))}
    </div>
  );
}
