import { Widget } from "./widget";

export function WidgetList() {
  return (
    <div className="sticky inset-y-0 h-dvh w-72 p-4 flex-shrink-0 border-l hidden md:block">
      <div className="grid gap-4">
        {Array.from({ length: 4 }).map((_, i) => (
          // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
          <Widget key={i} />
        ))}
      </div>
    </div>
  );
}
