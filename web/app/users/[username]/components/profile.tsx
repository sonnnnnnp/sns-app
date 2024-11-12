/* eslint-disable @next/next/no-img-element */
"use client";

import { PostList } from "@/app/home/components/post-list";
import { ConfirmActionDialog } from "@/components/dialog/confirm-action-dialog";
import { MainCard } from "@/components/main-card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
} from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Separator } from "@/components/ui/separator";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import client from "@/lib/api";
import { components } from "@/lib/api/client";
import { convertNewlinesToBr } from "@/utils/text";
import {
  AtSignIcon,
  BanIcon,
  BellIcon,
  ChevronDownIcon,
  ChevronLeftIcon,
  EyeOffIcon,
  InfoIcon,
  Loader2,
  MailIcon,
  MinusIcon,
  PlusIcon,
  RefreshCwOffIcon,
  Share2Icon,
  SlidersHorizontal,
} from "lucide-react";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

export function Profile() {
  const router = useRouter();
  const pathParams = useParams<{ username: string }>();

  const [user, setUser] = useState<components["schemas"]["User"] | null>(null);
  const [posts, setPosts] = useState<components["schemas"]["Post"][]>([]);

  const [tabValue, setTabValue] = useState("posts");
  const [isFollowing, setIsFollowing] = useState(false);
  const [isBlockingUser, setIsBlockingUser] = useState(false);
  const [isBlockedByUser, setIsBlockedByUser] = useState(false);
  const [actionDialogOpen, setActionDialogOpen] = useState(false);

  const handleFetchProfile = useCallback(async () => {
    const user = await client.GET("/users/{name}", {
      params: { path: { name: pathParams.username } },
    });
    if (!user.data?.ok) {
      return console.error("error fetching user");
    }

    setUser(user.data.data);
    setIsFollowing(user.data.data.social_connection?.following ?? false);
    setIsBlockingUser(user.data.data.block_status?.blocking ?? false);

    const timeline = await client.GET("/timeline", {
      params: { query: { user_id: user.data.data.id } },
    });
    if (!timeline.data?.ok) {
      return console.error("error fetching timeline");
    }

    setPosts(timeline.data.data.posts);
  }, [pathParams.username]);

  useEffect(() => {
    handleFetchProfile();
  }, [handleFetchProfile]);

  const handleTabChange = (tab: string) => {
    setTabValue(tab);
  };

  const onFollowActionConfirm = async () => {
    if (!user) {
      return;
    }

    if (isFollowing) {
      const { data } = await client.DELETE("/users/{user_id}/follows", {
        params: {
          path: {
            user_id: user.id,
          },
        },
      });
      if (!data?.ok) {
        return;
      }
      setUser({
        ...user,
        social_engagement: {
          ...user.social_engagement,
          followers_count: (user.social_engagement &&
            user.social_engagement?.followers_count - 1) as number,
        } as components["schemas"]["User"]["social_engagement"], // 誰が見てもウンコード。形もウンコ。
      });
    } else {
      const { data } = await client.POST("/users/{user_id}/follows", {
        params: {
          path: {
            user_id: user.id,
          },
        },
      });
      if (!data?.ok) {
        return;
      }
      setUser({
        ...user,
        social_engagement: {
          ...user.social_engagement,
          followers_count: (user.social_engagement &&
            user.social_engagement?.followers_count + 1) as number,
        } as components["schemas"]["User"]["social_engagement"],
      });
    }

    setIsFollowing(!isFollowing);
    setActionDialogOpen(false);
  };

  const handleBlockUser = async () => {
    const { data } = await client.POST("/users/{user_id}/blocks", {
      params: {
        path: {
          user_id: user!.id,
        },
      },
    });

    if (data?.ok) {
      return setIsBlockingUser(true);
    }
  };

  const handleUnblockUser = async () => {
    const { data } = await client.DELETE("/users/{user_id}/blocks", {
      params: {
        path: {
          user_id: user!.id,
        },
      },
    });

    if (data?.ok) {
      setIsBlockingUser(false);
      handleFetchProfile();
    }
  };

  return (
    <div>
      <MainCard>
        <Tabs
          defaultValue="posts"
          onValueChange={handleTabChange}
          className="w-full"
        >
          <div className="flex flex-col w-full">
            <div className="sticky top-0 z-10 rounded-none w-full bg-background/60 backdrop-blur border-b">
              <div className="flex justify-between items-center h-14 px-2 text-muted-foreground">
                <Button variant="ghost" onClick={() => router.back()}>
                  <ChevronLeftIcon className="h-5 w-5" />
                </Button>
                <Button variant="ghost">
                  <SlidersHorizontal className="h-5 w-5" />
                </Button>
              </div>
            </div>
            <div className="flex flex-col space-y-4 w-full p-4">
              <Card className="p-0 overflow-hidden">
                <div className="relative">
                  {user?.banner_image_url ? (
                    <img
                      src={user.banner_image_url}
                      alt="banner"
                      className="object-cover w-full h-52 sm:h-64"
                    />
                  ) : (
                    <div className="w-full bg-accent h-52 sm:h-64" />
                  )}
                  <div className="absolute top-4 right-4">
                    <div className="flex items-center space-x-1 rounded-md bg-secondary text-secondary-foreground">
                      {isBlockingUser ? (
                        <Button
                          variant="secondary"
                          className="px-3 shadow-none"
                          onClick={handleUnblockUser}
                        >
                          <span className="flex items-center">
                            {actionDialogOpen ? (
                              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                            ) : (
                              <MinusIcon className="mr-2 h-4 w-4" />
                            )}
                            ブロック解除
                          </span>
                        </Button>
                      ) : (
                        <div className="flex items-center space-x-1 rounded-md bg-secondary text-secondary-foreground">
                          <ConfirmActionDialog
                            open={actionDialogOpen}
                            message={
                              isFollowing
                                ? "このユーザーのフォローを解除しますか？"
                                : "このユーザーをフォローしますか？"
                            }
                            onConfirm={onFollowActionConfirm}
                            onCancel={() => setActionDialogOpen(false)}
                          />
                          <Button
                            variant="secondary"
                            className="px-3 shadow-none"
                            onClick={() => setActionDialogOpen(true)}
                          >
                            {isBlockingUser ? (
                              <span className="flex items-center">
                                {actionDialogOpen ? (
                                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                                ) : (
                                  <MinusIcon className="mr-2 h-4 w-4" />
                                )}
                                ブロック解除
                              </span>
                            ) : isFollowing ? (
                              <span className="flex items-center">
                                {actionDialogOpen ? (
                                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                                ) : (
                                  <MinusIcon className="mr-2 h-4 w-4" />
                                )}
                                フォロー解除
                              </span>
                            ) : (
                              <span className="flex items-center">
                                {actionDialogOpen ? (
                                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                                ) : (
                                  <PlusIcon className="mr-2 h-4 w-4" />
                                )}
                                フォロー
                              </span>
                            )}
                          </Button>
                          <Separator
                            orientation="vertical"
                            className="h-[20px]"
                          />
                          <DropdownMenu>
                            <DropdownMenuTrigger asChild>
                              <Button
                                variant="secondary"
                                className="px-2 shadow-none"
                              >
                                <ChevronDownIcon className="h-4 w-4 text-secondary-foreground" />
                              </Button>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent
                              align="end"
                              alignOffset={-5}
                              forceMount
                            >
                              <DropdownMenuItem>
                                <AtSignIcon className="mr-2 h-4 w-4" />{" "}
                                ユーザー名をコピー
                              </DropdownMenuItem>
                              <DropdownMenuItem>
                                <Share2Icon className="mr-2 h-4 w-4" />
                                プロフィールURLをコピー
                              </DropdownMenuItem>
                              <DropdownMenuSeparator />
                              <DropdownMenuItem>
                                <MailIcon className="mr-2 h-4 w-4" />{" "}
                                メッセージを送信
                              </DropdownMenuItem>
                              <DropdownMenuItem>
                                <BellIcon className="mr-2 h-4 w-4" /> 投稿を通知
                              </DropdownMenuItem>
                              <DropdownMenuSeparator />
                              <DropdownMenuItem>
                                <EyeOffIcon className="mr-2 h-4 w-4" /> ミュート
                              </DropdownMenuItem>
                              <DropdownMenuItem>
                                <RefreshCwOffIcon className="mr-2 h-4 w-4" />{" "}
                                リポストをミュート
                              </DropdownMenuItem>
                              <DropdownMenuItem onClick={handleBlockUser}>
                                <BanIcon className="mr-2 h-4 w-4" /> ブロック
                              </DropdownMenuItem>
                              <DropdownMenuSeparator />
                              <DropdownMenuItem>
                                <InfoIcon className="mr-2 h-4 w-4" /> 通報
                              </DropdownMenuItem>
                            </DropdownMenuContent>
                          </DropdownMenu>
                        </div>
                      )}
                    </div>
                  </div>
                  <span className="absolute -bottom-10 left-4">
                    <Avatar className="h-32 w-32">
                      {user?.avatar_image_url ? (
                        <AvatarImage
                          src={user.avatar_image_url}
                          className="object-cover"
                        />
                      ) : (
                        <AvatarImage src="/users/placeholder-profile.svg" />
                      )}
                      <AvatarFallback>
                        <Loader2 className="h-8 w-8 text-muted-foreground animate-spin" />
                      </AvatarFallback>
                    </Avatar>
                  </span>
                </div>
                <div className="grid gap-y-4 pt-12 px-6 mb-2">
                  <div className="flex flex-col">
                    <span className="text-2xl">{user?.nickname}</span>
                    <span className="text-muted-foreground">{`@${user?.name}`}</span>
                  </div>
                  <p>{convertNewlinesToBr(user?.biography ?? "")}</p>
                  <div className="flex gap-x-3 text-xs">
                    <Link
                      href={`/users/${pathParams.username}/following`}
                      className="flex gap-x-1 hover:underline"
                    >
                      <span className="font-bold">
                        {user?.social_engagement?.following_count}
                      </span>
                      <span className="text-muted-foreground">フォロー中</span>
                    </Link>
                    <Link
                      href={`/users/${pathParams.username}/followers`}
                      className="flex gap-x-1 hover:underline"
                    >
                      <span className="font-bold">
                        {user?.social_engagement?.followers_count}
                      </span>
                      <span className="text-muted-foreground">フォロワー</span>
                    </Link>
                  </div>
                </div>
                {!isBlockedByUser && (
                  <TabsList className="w-full h-12 p-0 bg-background justify-between">
                    <TabsTrigger
                      value="posts"
                      className="relative w-full h-full font-bold rounded-none hover:bg-accent data-[state=active]:hover:bg-accent data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-12 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    >
                      <span>投稿</span>
                    </TabsTrigger>
                    <TabsTrigger
                      value="replies"
                      className="relative w-full h-full font-bold rounded-none hover:bg-accent data-[state=active]:hover:bg-accent data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-12 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    >
                      <span>返信</span>
                    </TabsTrigger>
                    <TabsTrigger
                      value="calls"
                      className="relative w-full h-full font-bold rounded-none hover:bg-accent data-[state=active]:hover:bg-accent data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-12 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    >
                      <span>通話</span>
                    </TabsTrigger>
                    <TabsTrigger
                      value="media"
                      className="relative w-full h-full font-bold rounded-none hover:bg-accent data-[state=active]:hover:bg-accent data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-12 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    >
                      <span>メディア</span>
                    </TabsTrigger>
                    <TabsTrigger
                      value="likes"
                      className="relative w-full h-full font-bold rounded-none hover:bg-accent data-[state=active]:hover:bg-accent data-[state=active]:shadow-none data-[state=active]:after:content-[''] data-[state=active]:after:bg-primary data-[state=active]:after:h-[3px] data-[state=active]:after:w-12 data-[state=active]:after:absolute data-[state=active]:after:bottom-0 data-[state=active]:after:rounded-md"
                    >
                      <span>いいね</span>
                    </TabsTrigger>
                  </TabsList>
                )}
              </Card>
              {isBlockingUser ? (
                <Card className="bg-transparent">
                  <CardHeader className="pb-2 text-2xl font-bold">
                    このユーザーをブロック中です 😢
                  </CardHeader>
                  <CardDescription className="px-6 pb-4">
                    {`このユーザーの投稿を表示しますか？投稿を見るだけなら @${user?.name} さんのブロックを解除しなくても見ることができます。`}
                  </CardDescription>
                  <CardContent>
                    <Button
                      className="w-full md:w-28"
                      onClick={() => setIsBlockingUser(false)}
                    >
                      投稿を表示
                    </Button>
                  </CardContent>
                </Card>
              ) : (
                <div>
                  {/* pinned posts (limit: 3..5) */}
                  <Card>
                    <PostList posts={posts} />
                  </Card>
                  {/* normal posts ordered by create_at desc with tabs contains
                posts, reply and media */}
                  <Card>
                    <PostList posts={posts} />
                  </Card>
                </div>
              )}
            </div>
          </div>
        </Tabs>
      </MainCard>
    </div>
  );
}
