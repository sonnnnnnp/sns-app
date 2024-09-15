import ChatListItem from "./ChatListItem";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Separator } from "@/components/ui/separator";

const ChatList = () => {
  return (
    <div className="flex flex-col h-full">
      <CardTitle className="m-4 text-xl">メッセージ</CardTitle>
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
