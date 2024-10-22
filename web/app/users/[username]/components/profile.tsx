/* eslint-disable @next/next/no-img-element */
"use client";

import { PostList } from "@/app/home/components/post-list";
import { MainCard } from "@/components/main-card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Separator } from "@/components/ui/separator";
import client from "@/lib/api";
import { components } from "@/lib/api/client";
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
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { ConfirmActionDialog } from "@/components/dialog/confirm-action-dialog";
import { useRouter } from "next/navigation";

export function Profile() {
  const router = useRouter();
  const pathParams = useParams<{ username: string }>();

  const [user, setUser] = useState<components["schemas"]["User"] | null>(null);
  const [posts, setPosts] = useState<components["schemas"]["Post"][]>([]);

  const [isFollowing, setIsFollowing] = useState(false);
  const [actionDialogOpen, setActionDialogOpen] = useState(false);

  useEffect(() => {
    const fetchProfile = async () => {
      const user = await client.GET("/users", {
        params: {
          query: {
            name: pathParams.username,
          },
        },
      });

      if (!user.data?.ok) {
        return console.error("error fetching user");
      }

      setUser(user.data.data);
      setIsFollowing(user.data.data.social_context?.following ?? false);

      const timeline = await client.GET("/timeline", {
        params: {
          query: {
            user_id: user.data.data.id,
          },
        },
      });

      if (!timeline.data?.ok) {
        return console.error("error fetching timeline");
      }

      setPosts(timeline.data.data.posts);
    };

    fetchProfile();
  }, [pathParams.username]);

  const onFollowActionConfirm = async () => {
    if (!user) {
      return;
    }

    if (isFollowing) {
      const { data } = await client.POST("/users/following/delete", {
        body: {
          user_id: user.id,
        },
      });
      if (!data?.ok) {
        return;
      }
    } else {
      const { data } = await client.POST("/users/following/create", {
        body: {
          user_id: user.id,
        },
      });
      if (!data?.ok) {
        return;
      }
    }

    setIsFollowing(!isFollowing);
    setActionDialogOpen(false);
  };

  return (
    <div>
      <MainCard>
        <div className="flex flex-col w-full">
          <div className="sticky top-0 z-10 rounded-none w-full bg-transparent backdrop-blur border-b">
            <div className="flex justify-between items-center h-14 px-2 text-muted-foreground">
              <Button variant="ghost" onClick={() => router.back()}>
                <ChevronLeftIcon className="h-5 w-5" />
              </Button>
              <Button variant="ghost">
                <SlidersHorizontal className="h-5 w-5" />
              </Button>
            </div>
          </div>
          <div className="flex flex-col space-y-3 w-full p-3">
            <Card className="p-0 overflow-hidden">
              <div className="relative">
                <img
                  src={
                    "https://i.pinimg.com/enabled_lo/564x/06/57/90/065790b2ab543db53a960ea21f5227f4.jpg"
                  }
                  alt="banner"
                  className="object-cover w-full h-64"
                />
                <div className="absolute top-4 right-4">
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
                      {isFollowing ? (
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
                    <Separator orientation="vertical" className="h-[20px]" />
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
                          <MailIcon className="mr-2 h-4 w-4" /> メッセージを送信
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
                        <DropdownMenuItem>
                          <BanIcon className="mr-2 h-4 w-4" /> ブロック
                        </DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem>
                          <InfoIcon className="mr-2 h-4 w-4" /> 通報
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </div>
                </div>
                <span className="absolute top-40 left-4">
                  <Avatar className="h-32 w-32">
                    {user?.avatar_image_url ? (
                      <AvatarImage src={user.avatar_image_url} />
                    ) : (
                      <AvatarImage src="/users/placeholder-profile.svg" />
                    )}
                    <AvatarFallback>
                      <Loader2 className="h-8 w-8 text-muted-foreground animate-spin" />
                    </AvatarFallback>
                  </Avatar>
                </span>
              </div>
              <div className="grid gap-y-4 pt-12 p-6 border-b">
                <div className="flex flex-col">
                  <span className="text-2xl">{user?.nickname}</span>
                  <span className="text-muted-foreground">{`@${user?.name}`}</span>
                </div>
                <p>{user?.biography}</p>
              </div>
              <div className="flex justify-center min-h-14 text-muted-foreground">
                <div className="flex justify-between w-3/4">
                  <Link
                    href={"#"}
                    className="flex flex-col justify-center items-center"
                  >
                    <span>32</span>
                    <span className="text-xs">投稿</span>
                  </Link>
                  <Link
                    href={`/users/${pathParams.username}/reviews`}
                    className="flex flex-col justify-center items-center"
                  >
                    <span>4</span>
                    <span className="text-xs">レビュー</span>
                  </Link>
                  <Link
                    href={`/users/${pathParams.username}/following`}
                    className="flex flex-col justify-center items-center"
                  >
                    <span>11</span>
                    <span className="text-xs">フォロー</span>
                  </Link>
                  <Link
                    href={`/users/${pathParams.username}/followers`}
                    className="flex flex-col justify-center items-center"
                  >
                    <span>232</span>
                    <span className="text-xs">フォロワー</span>
                  </Link>
                </div>
              </div>
            </Card>
            {/* pinned posts (limit: 3..5) */}
            <Card>
              <PostList posts={posts} />
            </Card>
            {/* normal posts ordered by create_at desc with tabs contains posts, reply and media */}
            <Card>
              <PostList posts={posts} />
            </Card>
          </div>
        </div>
      </MainCard>
    </div>
  );
}
