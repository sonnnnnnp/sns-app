import { cn } from "@nextui-org/react";

export function PostContentBody({ className }: { className?: string }) {
  return (
    <p className={cn("break-all", className)}>ここに投稿本文が入ります。</p>
  );
}
