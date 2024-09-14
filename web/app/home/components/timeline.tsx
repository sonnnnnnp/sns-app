import { Post } from "./post";
import { Tabs } from "./tabs";
import { TimelineTab } from "./timeline-tab";

export function Timeline() {
  return (
    <div className="mx-3">
      <TimelineTab />
      <div className="grid gap-y-3 pt-[68px]">
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
      </div>
    </div>
  );
}
