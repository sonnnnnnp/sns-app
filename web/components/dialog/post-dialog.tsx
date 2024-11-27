import * as VisuallyHidden from "@radix-ui/react-visually-hidden";

import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Images, Pencil, Type } from "lucide-react";
import { Avatar, AvatarImage } from "../ui/avatar";
import { Button } from "../ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import { Separator } from "../ui/separator";
import { Textarea } from "../ui/textarea";

type Props = {
  open: boolean;
  postContent: string;
  postContentValid: boolean;
  onCloseAction: () => void;
  onPostContentChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  handleDraftPost: () => void;
  handleCreatePost: () => void;
};

export const PostDialog = ({
  open,
  postContent,
  postContentValid,
  onCloseAction,
  onPostContentChange,
  handleDraftPost,
  handleCreatePost,
}: Props) => {
  return (
    <Dialog open={open}>
      <DialogContent
        hideCloseButton={true}
        onInteractOutside={() => onCloseAction()}
        onEscapeKeyDown={() => onCloseAction()}
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
              className="min-h-28 resize-none border-0 shadow-none focus-visible:ring-0 focus-visible:ring-offset-0"
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
  );
};
