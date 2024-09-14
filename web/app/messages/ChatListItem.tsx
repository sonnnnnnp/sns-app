import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

const ChatListItem = () => {
  return (
    <div className="m-1.5">
      <div className="flex">
        <div>
          <Avatar>
            <AvatarImage src="/kaworu_icon.jpg" />
            <AvatarFallback>xxx</AvatarFallback>
          </Avatar>
        </div>
        <div>
          <div>
            <span>カヲル</span>
            <span>@kaworu</span>
          </div>
          <div>それってあなたの感想ですよね</div>
        </div>
      </div>
    </div>
  );
};
export default ChatListItem;
