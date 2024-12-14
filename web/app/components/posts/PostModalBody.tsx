import { Avatar, Divider, Link, ModalBody, Textarea } from "@nextui-org/react";
import { PostContentBody } from "./PostContentBody";
import { PostContentHeader } from "./PostContentHeader";

type Post = {
  text: string;
};

export function PostModalBody({
  replyToPost,
  repost,
}: {
  replyToPost?: Post;
  repost?: Post;
}) {
  return (
    <ModalBody className="min-h-32 p-4 gap-0">
      {replyToPost && (
        <div className="flex gap-4">
          <div>
            <Avatar
              src="https://i.pravatar.cc/150?u=a04258114e29026702d"
              classNames={{ base: "flex-shrink-0" }}
            />
            <Divider orientation="vertical" className="mx-auto w-[2px]" />
          </div>
          <div>
            <PostContentHeader />
            <PostContentBody className="text-sm" />
            <span className="flex gap-1 h-12 items-center text-sm text-foreground-400">
              <span>返信先:</span>
              <Link>@username</Link>
            </span>
          </div>
        </div>
      )}
      <div className="flex grid-2">
        <Avatar
          src="https://i.pravatar.cc/150?u=a04258114e29026702d"
          classNames={{ base: "mt-1 flex-shrink-0" }}
        />
        <div className="grid gap-2 w-full">
          <Textarea
            autoFocus
            minRows={1}
            placeholder={
              replyToPost
                ? "@username に返信する"
                : repost
                  ? "コメントを入力する"
                  : "今日あったこと、興味のあること、なんでも気軽につぶやいてみよう！"
            }
            classNames={{
              inputWrapper:
                "bg-transparent shadow-none group-data-[focus=true]:bg-default-0 group-data-[focus-visible=true]:ring-0 group-data-[focus-visible=true]:ring-offset-0",
              input: "text-medium",
            }}
          />
          {repost && (
            <div className="flex gap-4 ml-2 px-3 py-2 border rounded-xl">
              <div>
                <PostContentHeader showAvatar />
                <PostContentBody className="text-sm" />
              </div>
            </div>
          )}
        </div>
      </div>
    </ModalBody>
  );
}
