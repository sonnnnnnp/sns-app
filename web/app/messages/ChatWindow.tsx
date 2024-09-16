import { Separator } from "@/components/ui/separator";
import { ChatMessageList } from "@/components/ui/chat/chat-message-list";
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
      <div className="flex flex-col flex-1 max-w-[480px]">
        <div className="flex-1 overflow-y-auto">
          <ChatMessageList>
            {[...Array(30)].map((_, i) => (
              <ChatBubble>
                <ChatBubbleAvatar />
                <ChatBubbleMessage>
                  Message and other content here
                </ChatBubbleMessage>
              </ChatBubble>
            ))}
            <ChatBubble className="max-w-[80%]">
              <ChatBubbleAvatar />
              <ChatBubbleMessage>
                この問題を解決するには、フレックスボックスレイアウトを使用して水平方向にコンポーネントを配置する必要があります。以下のように
                ChatWindow コンポーネントを修正することをお勧めします：
              </ChatBubbleMessage>
            </ChatBubble>
          </ChatMessageList>
        </div>
        <div>
          <Separator />
          <div>
            <ChatInput
              placeholder="Type your message here..."
              className="min-h-0"
            />
            <Button size="sm" className="ml-auto mt-2 gap-1.5">
              Send Message
              {/* <CornerDownLeft className="size-3.5" /> */}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};
export default ChatWindow;
