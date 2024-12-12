import { Tab, Tabs } from "@nextui-org/react";
import { PostTimeline } from "~/components/timeline/post-timeline";

export default function UserContent() {
  return (
    <Tabs
      fullWidth
      variant="underlined"
      classNames={{
        base: "sticky top-[55px] bg-background border-b z-10",
        tabList: "p-0",
        tab: "h-12",
        cursor: "w-[65%]",
        panel: "p-0",
      }}
    >
      <Tab key="posts" title="投稿">
        <PostTimeline />
      </Tab>
      <Tab key="replies" title="返信">
        <PostTimeline />
      </Tab>
      <Tab key="calls" title="通話">
        <PostTimeline />
      </Tab>
      <Tab key="media" title="メディア">
        <PostTimeline />
      </Tab>
      <Tab key="favorites" title="いいね">
        <PostTimeline />
      </Tab>
    </Tabs>
  );
}
