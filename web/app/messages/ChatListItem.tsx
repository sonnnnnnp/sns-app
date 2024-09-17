import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Separator } from "@/components/ui/separator";

const ChatListItem = () => {
  return (
    <>
      <div className="px-2 py-3" tabIndex={0}>
        <div className="flex">
          <div className="mx-2">
            <Avatar>
              <AvatarImage src="/kaworu_icon.jpg" />
              <AvatarFallback>xxx</AvatarFallback>
            </Avatar>
          </div>
          <div>
            <div>
              <span className="text-sm font-semibold leading-none">カヲル</span>
              <span className="text-sm text-muted-foreground">@kaworu</span>
            </div>
            <div className="text-sm text-muted-foreground">
              それってあなたの感想ですよね
            </div>
          </div>
        </div>
      </div>
      <Separator />
    </>
  );
};
export default ChatListItem;
