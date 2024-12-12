import { Button } from "@nextui-org/react";
import {
  EllipsisVerticalIcon,
  HeartIcon,
  MessageCircleIcon,
  Repeat2Icon,
} from "lucide-react";

export function PostActionButton({
  icon,
  label,
  ariaLabel,
}: { icon: React.ReactNode; label?: string; ariaLabel: string }) {
  return (
    <span className="relative flex items-center text-xs text-foreground-400">
      <Button
        isIconOnly
        size="sm"
        variant="light"
        aria-label={ariaLabel}
        className="rounded-full"
      >
        {icon}
      </Button>
      {label && <span className="absolute left-7">{label}</span>}
    </span>
  );
}

export function PostActions() {
  return (
    <div className="flex justify-between ml-10">
      <PostActionButton
        label="4"
        ariaLabel="reply"
        icon={<MessageCircleIcon className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        ariaLabel="repost"
        icon={<Repeat2Icon className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        label="18"
        ariaLabel="favorite"
        icon={<HeartIcon className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        ariaLabel="options"
        icon={<EllipsisVerticalIcon className="h-4 w-4 text-foreground-400" />}
      />
    </div>
  );
}
