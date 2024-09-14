import { Post } from "./post";
import { Tabs } from "./tabs";
import { TimelineTab } from "./timeline-tab";

export function Timeline() {
  return (
    <div className="px-3">
      <TimelineTab />
      <div className="grid gap-y-3 pt-[68px]">
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
        <Post
          username="kaworu"
          display_name="カヲル"
          avatar_image_url="/kaworu_icon.jpg"
          content="ユーザーがアクセスできるかは別として貴重なデータとして保存はしておきたい"
          created_at={1234567}
        />
      </div>
    </div>
  );
}
