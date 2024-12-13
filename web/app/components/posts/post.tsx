import { useDisclosure } from "@nextui-org/react";
import { PostActions } from "./post-actions";
import { PostContent } from "./post-content";
import { PostModal } from "./post-modal";

export function Post() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <div className="flex flex-col p-4 pb-1 border-b last:border-none">
      <PostContent />
      <PostActions onPressCreateReply={onOpen} />
      <PostModal
        replyToPost={{ text: "返信先の投稿本文" }}
        isOpen={isOpen}
        onOpenChange={onOpenChange}
      />
    </div>
  );
}
