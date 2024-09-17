import { Separator } from "@/components/ui/separator";
import { ScrollableChatMessageList } from "@/components/ui/chat/scrollable-chat-message-list";
import {
  ChatBubble,
  ChatBubbleAvatar,
  ChatBubbleMessage,
} from "@/components/ui/chat/chat-bubble";
import { ChatInput } from "@/components/ui/chat/chat-input";
import { Button } from "@/components/ui/button";
import { Info } from "lucide-react";
import { ChatHeader, ChatHeaderTitle, ChatHeaderButtons } from "./ChatHeader";
import React from "react";

const ChatWindow = () => {
  return (
    <div className="hidden md:flex h-full">
      <Separator orientation="vertical" className="h-full" />
      <div className="flex flex-col w-[500px]">
        <ChatHeader>
          <ChatHeaderTitle>カヲル</ChatHeaderTitle>
          <ChatHeaderButtons>
            <Info />
          </ChatHeaderButtons>
        </ChatHeader>
        <Separator />
        <ScrollableChatMessageList>
          {[...Array(30)].map((_, i) => (
            <React.Fragment key={i}>
              <ChatBubble>
                <ChatBubbleAvatar src="/kaworu_icon.jpg" />
                <ChatBubbleMessage>
                  なんかプログラミングってやつすれば色々作れるらしいぜ
                </ChatBubbleMessage>
              </ChatBubble>
              <ChatBubble className="max-w-[80%]" variant={"sent"}>
                <ChatBubbleAvatar src="/kaworu_icon.jpg" />
                <ChatBubbleMessage variant={"sent"}>まじ？</ChatBubbleMessage>
              </ChatBubble>
            </React.Fragment>
          ))}
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
