import { useDisclosure } from "@nextui-org/react";
import { PostActions } from "./PostActions";
import { PostModal } from "./PostModal";
import { PostContent } from "./PostContent";

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
