import ChatList from "./ChatList";
import ChatWindow from "./ChatWindow";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";

const Messages = () => {
  return (
    <div className="h-dvh py-7">
      <Card className="h-full flex">
        <ChatList />
        <Separator orientation="vertical" />
        <ChatWindow />
      </Card>
    </div>
  );
};

export default Messages;
