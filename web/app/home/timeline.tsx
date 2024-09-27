import { PlusCircle } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { PostList } from "./components/post-list";

export const iframeHeight = "938px";

export const containerClassName = "w-full h-full";

export default function Timeline() {
  return (
    <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8 max-w-[780px]">
      <Tabs defaultValue="following">
        <div className="flex items-center">
          <TabsList>
            <TabsTrigger value="following">フォロー中</TabsTrigger>
            <TabsTrigger value="public">オープン</TabsTrigger>
            <TabsTrigger value="live">配信</TabsTrigger>
          </TabsList>
          <div className="ml-auto flex items-center">
            <Button size="sm" className="h-8 gap-1">
              <PlusCircle className="h-3.5 w-3.5" />
              <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
                投稿する
              </span>
            </Button>
          </div>
        </div>
        <TabsContent value="following">
          <PostList />
        </TabsContent>
        <TabsContent value="public">
          <PostList />
        </TabsContent>
        <TabsContent value="live">
          <PostList />
        </TabsContent>
      </Tabs>
    </main>
  );
}
