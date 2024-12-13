import { Tab, Tabs } from "@nextui-org/react";
import { Timeline } from "~/components/timeline/timeline";

export default function Home() {
  return (
    <Tabs
      aria-label="timeline-tab"
      variant="underlined"
      fullWidth
      classNames={{
        base: "sticky top-0 backdrop-blur-md border-b z-[99]",
        tabList: "p-0",
        tab: "h-14 font-bold",
        cursor: "w-20",
        panel: "p-0",
      }}
    >
      <Tab key="following" title="フォロー中">
        <Timeline />
      </Tab>
      <Tab key="public" title="オープン">
        <Timeline />
      </Tab>
    </Tabs>
  );
}
