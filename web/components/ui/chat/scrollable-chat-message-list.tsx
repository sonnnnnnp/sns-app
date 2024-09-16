import * as React from "react";
import { cn } from "@/lib/utils";
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";

interface ScrollableChatMessageListProps
  extends React.HTMLAttributes<HTMLDivElement> {}

const ScrollableChatMessageList = React.forwardRef<
  HTMLDivElement,
  ScrollableChatMessageListProps
>(({ className, children, ...props }, ref) => (
  <ScrollArea
    className={cn("w-full h-full", className)}
    {...(props as React.ComponentPropsWithoutRef<typeof ScrollArea>)}
  >
    <div className="flex flex-col w-full p-4 gap-6" ref={ref}>
      {children}
    </div>
    <ScrollBar />
  </ScrollArea>
));

ScrollableChatMessageList.displayName = "ScrollableChatMessageList";

export { ScrollableChatMessageList };
