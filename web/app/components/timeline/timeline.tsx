import { Call } from "../calls/call";
import { Post } from "../posts/post";

export function Timeline() {
  return (
    <div>
      {/* 通話のタイムライン */}
      <div className="w-full px-4 py-5 border-b overflow-x-auto">
        <div className="flex gap-4">
          {Array.from({ length: 8 }).map((_, i) => (
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
    </div>
  );
}
