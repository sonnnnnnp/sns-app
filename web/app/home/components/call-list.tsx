import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
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
  return (
    <div>
      <ScrollArea className="border-b">
        {calls.length > 0 ? (
          <div className="flex w-max space-x-5 p-4 pb-5">
            {calls.map((call, i) => (
              <Call key={i} call={call} />
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
