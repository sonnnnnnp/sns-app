import {
  Button,
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@nextui-org/react";
import React from "react";
import {
  SolarChatRoundLinear,
  SolarHeartLinear,
  SolarMenuDotsBold,
  SolarRepeatLinear,
} from "../icons";
import { PostOptionsList } from "./PostOptionsList";

export function PostActionButton({
  icon,
  label,
  ariaLabel,
  onPress,
}: {
  icon: React.ReactNode;
  label: string;
  ariaLabel: string;
  onPress?: () => void;
}) {
  return (
    <span className="relative flex items-center text-xs text-foreground-400">
      <Button
        isIconOnly
        size="sm"
        variant="light"
        aria-label={ariaLabel}
        className="rounded-full"
        onPress={onPress}
      >
        {icon}
      </Button>
      <span className="absolute left-7">{label}</span>
    </span>
  );
}

export function PostActions({
  onPressCreateReply,
}: { onPressCreateReply: () => void }) {
  const [isOptionsPopoverOpen, setIsOptionsPopoverOpen] = React.useState(false);

  return (
    <div className="flex justify-between ml-10">
      <PostActionButton
        label="4"
        ariaLabel="reply"
        icon={<SolarChatRoundLinear className="h-4 w-4 text-foreground-400" />}
        onPress={onPressCreateReply}
      />
      <PostActionButton
        label=""
        ariaLabel="repost"
        icon={<SolarRepeatLinear className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        label="18"
        ariaLabel="favorite"
        icon={<SolarHeartLinear className="h-4 w-4 text-foreground-400" />}
      />
      <Popover
        placement="left-end"
        isOpen={isOptionsPopoverOpen}
        onOpenChange={(open) => setIsOptionsPopoverOpen(open)}
      >
        <PopoverTrigger>
          <Button
            isIconOnly
            size="sm"
            variant="light"
            aria-label="options"
            className="rounded-full"
          >
            <SolarMenuDotsBold className="h-4 w-4 rotate-90 text-foreground-400" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="min-w-[140px] border-small p-0">
          <PostOptionsList onSelect={() => setIsOptionsPopoverOpen(false)} />
        </PopoverContent>
      </Popover>
    </div>
  );
}
