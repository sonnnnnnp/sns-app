import { Tabs, TabsContent, TabsList, TabsTrigger } from "./tabs";
import { PostList } from "./post-list";

export function Timeline() {
  useEffect(() => {
    (async () => {
      await client.GET("/timeline");
    })();
  }, []);
  return (
    <Tabs className="z-50 md:w-[580px]" defaultValue="following">
      <div>
        <TabsList className="h-14 w-full sticky top-0 z-10">
          <TabsTrigger value="following">フォロー中</TabsTrigger>
          <TabsTrigger value="recommendations">おすすめ</TabsTrigger>
        </TabsList>
        <TabsContent value="following">
          <PostList />
        </TabsContent>
        <TabsContent value="recommendations">
          <PostList />
        </TabsContent>
      </div>
    </Tabs>
  );
}
