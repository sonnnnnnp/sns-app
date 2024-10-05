"use client";

import { Bell, Home, MessagesSquare, Phone, Search, Users } from "lucide-react";
import { Button } from "@/components/ui/button";

export default function FooterNavigation() {
  return (
    <div className="fixed bottom-0 z-30 flex h-16 w-full items-center justify-between border-t bg-background px-6 sm:hidden">
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Home className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Search className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Phone className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Users className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <Bell className="h-6 w-6 text-muted-foreground" />
      </Button>
      <Button
        variant="ghost"
        size="icon"
        className="overflow-hidden rounded-full"
      >
        <MessagesSquare className="h-6 w-6 text-muted-foreground" />
      </Button>
    </div>
  );
}
