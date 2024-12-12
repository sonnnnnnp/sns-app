import { PostActions } from "./post-actions";
import { PostContent } from "./post-content";

export function Post() {
  return (
    <div className="flex flex-col p-4 pb-1 border-b last:border-none">
      <PostContent />
      <PostActions />
    </div>
  );
}
