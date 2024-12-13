import { Button } from "@nextui-org/react";
import {
  SolarChatRoundLinear,
  SolarHeartLinear,
  SolarMenuDotsBold,
  SolarRepeatLinear,
} from "../icons";

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
        icon={<SolarChatRoundLinear className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        ariaLabel="repost"
        icon={<SolarRepeatLinear className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        label="18"
        ariaLabel="favorite"
        icon={<SolarHeartLinear className="h-4 w-4 text-foreground-400" />}
      />
      <PostActionButton
        ariaLabel="options"
        icon={
          <SolarMenuDotsBold className="h-4 w-4 rotate-90 text-foreground-400" />
        }
      />
    </div>
  );
}
