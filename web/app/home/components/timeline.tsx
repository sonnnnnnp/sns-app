"use client";

import { Filter, Images, Pencil, Type } from "lucide-react";

import * as VisuallyHidden from "@radix-ui/react-visually-hidden";

import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { PostList } from "./post-list";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Textarea } from "@/components/ui/textarea";
import { TooltipProvider } from "@radix-ui/react-tooltip";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Separator } from "@/components/ui/separator";
import { ChangeEvent, useEffect, useState } from "react";
import client from "@/lib/api";
import { toast } from "sonner";

import twitterText from "twitter-text";
import { components } from "@/lib/api/client";

export const iframeHeight = "938px";

export const containerClassName = "w-full h-full";

export default function Timeline() {
  const [postDialogOpen, setPostDialogOpen] = useState(false);
  const [postContentValid, setPostContentValid] = useState(false);
  const [postContent, setPostContent] = useState("");

  const [posts, setPosts] = useState<
    components["schemas"]["Timeline"]["posts"]
  >([]);

  useEffect(() => {
    const loadTimeline = async () => {
      fetchTimeline("following");
    };

    loadTimeline();
  }, []);

  const fetchTimeline = async (value: string) => {
    const { data } = await client.GET("/timeline");
    if (!data?.ok) {
      return console.error("error fetching timeline");
    }

    console.log(value);

    setPosts(data?.data.posts ?? []);
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
    if (postContent.length <= 0) {
      toast("エラー", {
        description: "本文を入力してください",
      });
      return;
    }

    const { data } = await client.POST("/posts/create", {
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
    <div className="max-w-[780px]">
      <Tabs defaultValue="following" onValueChange={fetchTimeline}>
        <div className="flex items-center">
          <TabsList>
            <TabsTrigger value="following">フォロー中</TabsTrigger>
            <TabsTrigger value="public">オープン</TabsTrigger>
            <TabsTrigger value="live">配信</TabsTrigger>
          </TabsList>
          <div className="ml-auto flex items-center">
            <Button size="sm" className="h-8 gap-1">
              <Filter className="h-3.5 w-3.5" />
              <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
                絞り込み
              </span>
            </Button>
          </div>
        </div>
        <TabsContent value="following">
          <PostList posts={posts} />
        </TabsContent>
        <TabsContent value="public">
          <PostList posts={posts} />
        </TabsContent>
        <TabsContent value="live">
          <PostList posts={posts} />
        </TabsContent>
      </Tabs>
      <div className="fixed bottom-[12%] right-[10%]">
        <Button
          size="icon"
          className="h-14 w-14 overflow-hidden rounded-full"
          onClick={() => {
            setPostDialogOpen(true);
          }}
        >
          <Pencil className=" h-6 w-6" />
        </Button>
        <Dialog open={postDialogOpen}>
          <DialogContent
            hideCloseButton={true}
            onInteractOutside={() => setPostDialogOpen(false)}
            onEscapeKeyDown={() => setPostDialogOpen(false)}
            className="pb-3"
          >
            <VisuallyHidden.Root>
              <DialogHeader>
                <DialogTitle />
                <DialogDescription />
              </DialogHeader>
            </VisuallyHidden.Root>
            <div className="flex">
              <Avatar className="h-9 w-9 mt-0.5">
                <AvatarImage src="/users/placeholder-profile.svg" />
              </Avatar>
              <div className="relative w-full pr-2 overflow-hidden bg-background">
                <Textarea
                  placeholder="なんでも気軽につぶやいてみよう！"
                  className="min-h-28 resize-none border-0 shadow-none focus-visible:ring-0"
                  onChange={onPostContentChange}
                  value={postContent}
                />
              </div>
            </div>
            <Separator />
            <div className="flex items-center pt-0">
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button variant="ghost" size="icon">
                      <Images className="size-4" />
                      <span className="sr-only">メディア</span>
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent side="top">メディア</TooltipContent>
                </Tooltip>
              </TooltipProvider>
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button variant="ghost" size="icon">
                      <Type className="size-4" />
                      <span className="sr-only">フォント</span>
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent side="top">フォント</TooltipContent>
                </Tooltip>
              </TooltipProvider>
              <Button
                variant="link"
                className="ml-auto gap-1.5"
                onClick={handleDraftPost}
              >
                下書き
              </Button>
              <Button
                size="sm"
                className="ml-2 gap-1.5"
                onClick={handleCreatePost}
                disabled={!postContentValid}
              >
                投稿する
                <Pencil className="size-3.5" />
              </Button>
            </div>
          </DialogContent>
        </Dialog>
      </div>
    </div>
  );
}
