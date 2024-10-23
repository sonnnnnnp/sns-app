import { Card, CardDescription } from "../ui/card";

export function Widget() {
  return (
    <aside className="sticky inset-y-0 w-80 h-dvh p-4 hidden overflow-scroll lg:block">
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
    </aside>
  );
}
