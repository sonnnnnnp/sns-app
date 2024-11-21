import { ReactNode } from "react";

interface ContainerProps {
  children: ReactNode;
}

const ChatHeader: React.FC<ContainerProps> = ({ children }) => {
  return <div className="flex items-center mx-4 min-h-[52px]">{children}</div>;
};

const ChatHeaderTitle: React.FC<ContainerProps> = ({ children }) => {
  return <h2 className="text-xl font-semibold">{children}</h2>;
};

const ChatHeaderButtons: React.FC<ContainerProps> = ({ children }) => {
  return <div className="ml-auto">{children}</div>;
};
export { ChatHeader, ChatHeaderTitle, ChatHeaderButtons };
