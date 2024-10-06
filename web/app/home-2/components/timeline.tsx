import { Tabs, TabsContent, TabsList, TabsTrigger } from "./tabs";
import { PostList } from "../../home/components/post-list";

export function Timeline() {
  return (
    <Tabs className="z-50 md:w-[580px]" defaultValue="following">
      <div>
        <TabsList className="h-14 w-full sticky top-0 z-10">
          <TabsTrigger value="following">フォロー中</TabsTrigger>
          <TabsTrigger value="recommendations">おすすめ</TabsTrigger>
        </TabsList>
        <TabsContent value="following">
          <PostList posts={[]} />
        </TabsContent>
        <TabsContent value="recommendations">
          <PostList posts={[]} />
        </TabsContent>
      </div>
    </Tabs>
  );
}
