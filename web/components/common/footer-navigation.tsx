import { Bell, Blocks, Home, MessagesSquare } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarImage } from "../ui/avatar";

export default function FooterNavigation() {
  return (
    <aside className="sticky bottom-0 z-30 flex items-center justify-between w-full px-6 h-16 border-t bg-background sm:hidden">
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Home className="h-6 w-6 text-primary" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Blocks className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <MessagesSquare className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Bell className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="outline"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Avatar className="h-9 w-9">
          <AvatarImage src="/users/placeholder-profile.svg" />
        </Avatar>
      </Button>
    </aside>
  );
}
