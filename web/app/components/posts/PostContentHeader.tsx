export function PostContentHeader() {
  return (
    <div className="flex gap-1.5 mb-1">
      <span className="font-bold">ユーザー名</span>
      <span className="text-foreground-400">@username</span>
      <span className="ml-auto text-xs text-foreground-400">3時間前</span>
    </div>
  );
}
