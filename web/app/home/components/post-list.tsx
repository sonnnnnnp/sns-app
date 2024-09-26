import { Post } from "./post";

export function PostList() {
  return (
    <div className="grid gap-y-3">
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
