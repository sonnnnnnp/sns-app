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

const Messages = () => {
  return (
    <div className="absolute top-0 left-0 h-dvh py-7">
      <Card className="h-full flex">
        <ChatList />
        <ChatWindow />
      </Card>
    </div>
  );
};

export default Messages;
