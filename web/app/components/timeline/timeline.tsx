import { Call } from "../calls/call";
import { Post } from "../posts/post";

export function Timeline() {
  return (
    <div className="w-full">
      {/* 通話のタイムライン */}
      <div className="w-full px-6 pt-5 pb-3 overflow-x-scroll">
        <div className="flex space-x-5">
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
    </div>
  );
}
