import { Post } from "./post";
import { Tabs } from "./tabs";
import { TimelineTab } from "./timeline-tab";

export function Timeline() {
  return (
    <div className="px-3">
      <TimelineTab />
      <div className="grid gap-y-3 pt-[68px]">
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
    </div>
  );
}
