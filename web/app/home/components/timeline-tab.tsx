import { Tabs, TabsContent, TabsList, TabsTrigger } from "./tabs";

export function TimelineTab() {
  return (
    <Tabs className="fixed w-[580px] h-14 z-50" defaultValue="recommendations">
      <TabsList>
        <TabsTrigger value="recommendations">おすすめ</TabsTrigger>
        <TabsTrigger value="following">フォロー中</TabsTrigger>
      </TabsList>
    </Tabs>
  );
}
