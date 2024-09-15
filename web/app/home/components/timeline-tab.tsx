import { Tabs, TabsContent, TabsList, TabsTrigger } from "./tabs";

export function TimelineTab() {
  return (
    <div className="fixed w-[calc(100%_-_1.5rem)]">
      <Tabs className="md:w-[580px] h-14 z-50" defaultValue="recommendations">
        <TabsList>
          <TabsTrigger value="recommendations">おすすめ</TabsTrigger>
          <TabsTrigger value="following">フォロー中</TabsTrigger>
        </TabsList>
      </Tabs>
    </div>
  );
}