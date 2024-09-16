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
      <div className="flex-1 overflow-y-auto">
        {[...Array(30)].map((_, index) => (
          <ChatListItem key={index} />
        ))}
      </div>
    </div>
  );
};
export default ChatList;
