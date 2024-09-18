import ChatListItem from "./ChatListItem";
import { Separator } from "@/components/ui/separator";
import { SquarePen } from "lucide-react";
import { ScrollArea } from "@/components/ui/scroll-area";
import { ChatHeader, ChatHeaderTitle, ChatHeaderButtons } from "./ChatHeader";

const ChatList = () => {
  return (
    <div className="flex flex-col h-full w-[calc(100dvw-1rem)] md:w-80">
      <ChatHeader>
        <ChatHeaderTitle>メッセージ</ChatHeaderTitle>
        <ChatHeaderButtons>
          <SquarePen />
        </ChatHeaderButtons>
      </ChatHeader>
      <Separator />
      <ScrollArea className="flex-1">
        {[...Array(20)].map((_, index) => (
          <ChatListItem key={index} />
        ))}
      </ScrollArea>
    </div>
  );
};
export default ChatList;
