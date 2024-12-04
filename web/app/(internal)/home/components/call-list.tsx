import { CallTimelineCarousel } from "@/components/call-timeline-carousel";
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import React from "react";
import { Call } from "./call";

type CallUser = {
  id: string;
  name: string;
  avatar_url?: string;
};

export type CallData = {
  host: CallUser;
  content: string;
  id: string;
  type: "voice" | "video";
  participants: CallUser[];
};

type Props = {
  calls: CallData[];
};

export function CallList({ calls }: Props) {
  const [isCarouselOpen, setCarouselOpen] = React.useState(false);
  const [carouselIndex, setCarouselIndex] = React.useState(0);

  const showCallListCarousel = (index: number) => {
    setCarouselOpen(true);
    setCarouselIndex(index);
  };

  return (
    <div>
      <CallTimelineCarousel
        index={carouselIndex}
        open={isCarouselOpen}
        onOpenChange={() => setCarouselOpen(!isCarouselOpen)}
      />
      <ScrollArea className="p-4 pb-5 border-b">
        {calls.length > 0 ? (
          <div className="flex w-0 space-x-4">
            {calls.map((call, i) => (
              <a key={i} onClick={() => showCallListCarousel(i)}>
                <Call call={call} />
              </a>
            ))}
          </div>
        ) : (
          <p className="p-3 text-center text-gray-500">
            アクティブな通話は見つかりませんでした
          </p>
        )}
        <ScrollBar className="h-3" orientation="horizontal" />
      </ScrollArea>
    </div>
  );
}
