import { Skeleton } from "@/components/ui/skeleton";
import { Post } from "./post";
import { components } from "@/lib/api/client";
import { Separator } from "@/components/ui/separator";

type Props = {
  posts: components["schemas"]["Timeline"]["posts"];
};

export function PostList({ posts }: Props) {
  return (
    <div>
      {posts.length === 0
        ? Array.from({ length: 20 }).map((_, i) => (
            <div key={i} className="flex border-b pl-4 pr-3 pt-4 pb-6">
              <div className="mt-2 mr-2">
                <Skeleton className="h-10 w-10 rounded-full" />
              </div>
              <div className="flex-1 grid gap-y-2">
                <div className="flex items-center"></div>
                <div className="grid gap-y-1 space-y-4">
                  <div className="space-y-2">
                    <Skeleton className="h-4 w-[80%]" />
                    <Skeleton className="h-4 w-[30%]" />
                  </div>
                  <Skeleton className="h-4 w-[80%]" />
                </div>
              </div>
            </div>
          ))
        : posts.map((post, i) => (
            <div key={i}>
              <Post post={post} />
              {i < posts.length - 1 && <Separator />}
            </div>
          ))}
    </div>
  );
}
