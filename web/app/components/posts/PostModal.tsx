import { Modal, ModalContent } from "@nextui-org/react";
import { PostModalBody } from "./PostModalBody";
import { PostModalFooter } from "./PostModalFooter";
import { PostModalHeader } from "./PostModalHeader";

type Post = {
  text: string;
};

export function PostModal({
  isOpen,
  replyToPost,
  repost,
  onOpenChange,
}: {
  isOpen: boolean;
  replyToPost?: Post;
  repost?: Post;
  onOpenChange: (isOpen: boolean) => void;
}) {
  return (
    <Modal
      isOpen={isOpen}
      onOpenChange={onOpenChange}
      placement="center"
      size="lg"
      classNames={{ closeButton: "hidden" }}
      motionProps={{
        variants: {
          enter: {
            scale: 1,
            opacity: 1,
            transition: {
              duration: 0.1,
              ease: "easeOut",
            },
          },
          exit: {
            scale: 1.2,
            opacity: 0,
            transition: {
              duration: 0.2,
              ease: "easeIn",
            },
          },
        },
      }}
    >
      <ModalContent>
        {(onClose) => (
          <>
            <PostModalHeader onAction={onClose} isReply={!!replyToPost} />
            <PostModalBody replyToPost={replyToPost} repost={repost} />
            <PostModalFooter onAction={onClose} />
          </>
        )}
      </ModalContent>
    </Modal>
  );
}
