import { Post } from "./post";
import { components } from "@/lib/api/client";

type Props = {
  posts: components["schemas"]["Timeline"]["posts"];
};

export function PostList({ posts }: Props) {
  return (
    <div>
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
