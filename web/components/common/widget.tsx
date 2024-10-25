import { Card, CardDescription } from "../ui/card";
import { ScrollArea, ScrollBar } from "../ui/scroll-area";

export function Widget() {
  return (
    <aside className="sticky inset-y-0 h-dvh w-80 border-l hidden lg:block">
      <ScrollArea className="p-4 h-dvh">
        <div className="flex flex-col space-y-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <Card
              key={i}
              className="w-full h-[140px] flex items-center justify-center border"
            >
              <CardDescription>widget</CardDescription>
            </Card>
          ))}
        </div>
        <ScrollBar />
      </ScrollArea>
    </aside>
  );
}
