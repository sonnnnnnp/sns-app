import { Separator } from "@/components/ui/separator";
import { ScrollableChatMessageList } from "@/components/ui/chat/scrollable-chat-message-list";
import {
  ChatBubble,
  ChatBubbleAvatar,
  ChatBubbleMessage,
} from "@/components/ui/chat/chat-bubble";
import { ChatInput } from "@/components/ui/chat/chat-input";
import { Button } from "@/components/ui/button";

const ChatWindow = () => {
  return (
    <div className="hidden md:flex h-full">
      <Separator orientation="vertical" className="h-full" />
      <div className="flex flex-col w-[500px]">
        <ScrollableChatMessageList>
          {[...Array(30)].map((_, i) => (
            <ChatBubble key={i}>
              <ChatBubbleAvatar src="/kaworu_icon.jpg" />
              <ChatBubbleMessage>
                Message and other content here
              </ChatBubbleMessage>
            </ChatBubble>
          ))}
          <ChatBubble className="max-w-[80%]" variant={"sent"}>
            <ChatBubbleAvatar src="/kaworu_icon.jpg" />
            <ChatBubbleMessage variant={"sent"}>
              この問題を解決するには、フレックスボックスレイアウトを使用して水平方向にコンポーネントを配置する必要があります。以下のように
              ChatWindow コンポーネントを修正することをお勧めします：
            </ChatBubbleMessage>
          </ChatBubble>
        </ScrollableChatMessageList>
        {/* ~~~~~~~~~~ input ~~~~~~~~~~ */}
        <div>
          <Separator />
          <div className="flex items-center p-1">
            <ChatInput
              placeholder="Type your message here..."
              className="min-h-0"
            />
            <Button size="sm" className="ml-1">
              送信
            </Button>
          </div>
        </div>
        {/* ~~~~~~~~~~ input ~~~~~~~~~~ */}
      </div>
    </div>
  );
};

export default ChatWindow;
