"use client";

import { Pencil } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import client from "@/lib/api";
import { ChangeEvent, useEffect, useState } from "react";
import { toast } from "sonner";
import { PostList } from "./post-list";

import { PostDialog } from "@/components/dialog/post-dialog";
import { MainCard } from "@/components/main-card";
import { Card } from "@/components/ui/card";
import { components } from "@/lib/api/client";
import twitterText from "twitter-text";
import { CallData, CallList } from "./call-list";

export const iframeHeight = "938px";

export const containerClassName = "w-full h-full";

const calls: CallData[] = [
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/736x/e7/f4/35/e7f435a950216abefdaaed1112205a2c.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/736x/e7/f4/35/e7f435a950216abefdaaed1112205a2c.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/9d/40/be/9d40be148ccc55a4cfcb819c80422468.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/ca/e7/f2/cae7f23f24ba8e55b1c0d6bf2626dedf.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/ca/e7/f2/cae7f23f24ba8e55b1c0d6bf2626dedf.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/1e/d7/6b/1ed76b06b84ea5badca25933b8558f1a.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/ee/b4/55/eeb455cdfb13c5e47ed66379277212bb.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/96/35/3f/96353f3016dbeb412895aecad7ca0367.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/3f/ec/f7/3fecf7f1539948cc7b7c05c592f41a69.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/3f/ec/f7/3fecf7f1539948cc7b7c05c592f41a69.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/c2/4a/29/c24a29533c882d95bb8f2bd96b8f7813.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/736x/99/00/50/99005070f7697472a80739472c4d7517.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/originals/ef/7b/79/ef7b79862abd1877cf0e7f4ef375efde.gif",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/ef/7b/79/ef7b79862abd1877cf0e7f4ef375efde.gif",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/4b/a3/7d/4ba37db84be16aae53bb82779e1221c9.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/d7/f7/31/d7f731113dd15b79d935192c900e3beb.gif",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/d9/52/db/d952db68dce8bc132fe726a8b3ae0b83.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/c0/9d/9e/c09d9e14b66355edb8ff9ac1ac856116.gif",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/33/25/ba/3325ba8e900be9ce78e2f86825482d20.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/33/25/ba/3325ba8e900be9ce78e2f86825482d20.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/736x/e7/f4/35/e7f435a950216abefdaaed1112205a2c.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/736x/e7/f4/35/e7f435a950216abefdaaed1112205a2c.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/9d/40/be/9d40be148ccc55a4cfcb819c80422468.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/ca/e7/f2/cae7f23f24ba8e55b1c0d6bf2626dedf.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/ca/e7/f2/cae7f23f24ba8e55b1c0d6bf2626dedf.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/1e/d7/6b/1ed76b06b84ea5badca25933b8558f1a.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/ee/b4/55/eeb455cdfb13c5e47ed66379277212bb.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/96/35/3f/96353f3016dbeb412895aecad7ca0367.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/3f/ec/f7/3fecf7f1539948cc7b7c05c592f41a69.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/3f/ec/f7/3fecf7f1539948cc7b7c05c592f41a69.jpg",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/c2/4a/29/c24a29533c882d95bb8f2bd96b8f7813.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/736x/99/00/50/99005070f7697472a80739472c4d7517.jpg",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/originals/ef/7b/79/ef7b79862abd1877cf0e7f4ef375efde.gif",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/ef/7b/79/ef7b79862abd1877cf0e7f4ef375efde.gif",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/4b/a3/7d/4ba37db84be16aae53bb82779e1221c9.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/d7/f7/31/d7f731113dd15b79d935192c900e3beb.gif",
      },
      {
        id: "id3",
        name: "Charlie",
        avatar_url:
          "https://i.pinimg.com/474x/d9/52/db/d952db68dce8bc132fe726a8b3ae0b83.jpg",
      },
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/originals/c0/9d/9e/c09d9e14b66355edb8ff9ac1ac856116.gif",
      },
    ],
  },
  {
    id: "id1",
    content: "text",
    type: "voice",
    host: {
      id: "id1",
      name: "Alice",
      avatar_url:
        "https://i.pinimg.com/474x/33/25/ba/3325ba8e900be9ce78e2f86825482d20.jpg",
    },
    participants: [
      {
        id: "id2",
        name: "Bob",
        avatar_url:
          "https://i.pinimg.com/474x/33/25/ba/3325ba8e900be9ce78e2f86825482d20.jpg",
      },
    ],
  },
];

export function Timeline() {
  const [tabValue, setTabValue] = useState("following");
  const [postDialogOpen, setPostDialogOpen] = useState(false);
  const [postContentValid, setPostContentValid] = useState(false);
  const [postContent, setPostContent] = useState("");

  const [posts, setPosts] = useState<
    components["schemas"]["PostTimeline"]["posts"]
  >([]);

  useEffect(() => {
    onTabChange("following");
  }, []);

  const onTabChange = async (value: string) => {
    setTabValue(value);

    let timeline: components["schemas"]["PostTimeline"];

    if (value === "following") {
      const { data } = await client.GET("/timeline/following");
      if (!data?.ok) {
        return console.error("error fetching timeline");
      }
      timeline = data?.data;
    } else {
      const { data } = await client.GET("/timeline");
      if (!data?.ok) {
        return console.error("error fetching timeline");
      }
      timeline = data?.data;
    }

    setPosts(timeline.posts ?? []);
  };

  const onPostContentChange = async (e: ChangeEvent<HTMLTextAreaElement>) => {
    setPostContent(e.target.value);

    const parsedContent = twitterText.parseTweet(e.target.value);

    setPostContentValid(parsedContent.valid);
  };

  const handleDraftPost = async () => {
    setPostDialogOpen(false);
  };

  const handleCreatePost = async () => {
    const { data } = await client.POST("/posts", {
      body: {
        content: postContent,
      },
    });

    if (!data?.ok) {
      toast(`エラー: ${data?.code}`, {
        description: "投稿に失敗しました",
      });
      return;
    }

    // 投稿内容を先頭に挿入
    setPosts(() => {
      const newPosts = [...posts];
      newPosts.unshift(data?.data);
      return newPosts;
    });

    setPostDialogOpen(false);
  };

  return (
    <div>
      <MainCard>
        <div className="flex flex-col w-full">
          <Tabs defaultValue="following" onValueChange={onTabChange}>
            <TabsContent className="my-0" value={tabValue}>
              <div className="sticky top-0 z-10 rounded-none w-full">
                <TabsList className="h-14 p-0 rounded-none w-full bg-transparent backdrop-blur-md border-b">
                  <TabsTrigger
                    className="relative w-full h-full font-bold rounded-none bg-background/60 hover:bg-background/40 data-[state=active]:bg-background/60 data-[state=active]:hover:bg-background/40 data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-16 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    value="following"
                  >
                    フォロー中
                  </TabsTrigger>
                  <TabsTrigger
                    className="relative w-full h-full font-bold rounded-none bg-background/60 hover:bg-background/40 data-[state=active]:bg-background/60 data-[state=active]:hover:bg-background/40 data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-16 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    value="public"
                  >
                    オープン
                  </TabsTrigger>
                </TabsList>
              </div>
              <div className="flex flex-col space-y-4 w-full sm:p-4">
                <Card className="rounded-none sm:rounded-lg">
                  <CallList calls={calls} />
                  <PostList posts={posts} />
                </Card>
              </div>
            </TabsContent>
          </Tabs>
        </div>
      </MainCard>
      <div className="fixed bottom-[14%] right-[10%] sm:hidden">
        <Button
          size="icon"
          className="h-14 w-14 overflow-hidden rounded-full bg-primary/80 backdrop-blur-sm"
          onClick={() => {
            setPostDialogOpen(true);
          }}
        >
          <Pencil className="h-6 w-6" strokeWidth={1.5} />
        </Button>
        <PostDialog
          open={postDialogOpen}
          postContent={postContent}
          postContentValid={postContentValid}
          onCloseAction={() => setPostDialogOpen(false)}
          onPostContentChange={onPostContentChange}
          handleDraftPost={handleDraftPost}
          handleCreatePost={handleCreatePost}
        />
      </div>
    </div>
  );
}
