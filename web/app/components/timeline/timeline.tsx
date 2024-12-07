import { Card } from "@nextui-org/react";
import { Call } from "../calls/call";
import { Post } from "../posts/post";

export function Timeline() {
  return (
    <Card
      fullWidth
      shadow="none"
      classNames={{
        base: "bg-background rounded-none md:border md:rounded-large",
      }}
    >
      {/* 通話のタイムライン */}
      <div className="p-4 pb-2 overflow-x-scroll">
        <div className="flex space-x-4">
          {Array.from({ length: 2 }).map((_, i) => (
            // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
            <Call key={i} />
          ))}
        </div>
      </div>
      {/* 投稿のタイムライン */}
      {Array.from({ length: 20 }).map((_, i) => (
        // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
        <Post key={i} />
      ))}
    </Card>
  );
}
