import ChatListItem from "./ChatListItem";

const ChatList = () => {
  return (
    <div className="overflow-y-scroll">
      {[...Array(30)].map((_, index) => (
        <ChatListItem key={index} />
      ))}
    </div>
  );
};
export default ChatList;
