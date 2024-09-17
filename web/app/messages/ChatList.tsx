import ChatListItem from "./ChatListItem";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { SquarePen } from "lucide-react";
import { ScrollArea } from "@/components/ui/scroll-area";

const ChatList = () => {
  return (
    <div className="flex flex-col h-full w-[calc(100dvw-1rem)] md:w-80">
      <div className="flex items-center m-4">
        <CardTitle className="text-xl">メッセージ</CardTitle>
        <div className="ml-auto">
          <SquarePen />
        </div>
      </div>
      <Separator />
      <ScrollArea className="flex-1">
        {[...Array(30)].map((_, index) => (
          <ChatListItem key={index} />
        ))}
      </ScrollArea>
    </div>
  );
};
export default ChatList;
