import { Call } from "../calls/Call";

export function CallTimeline() {
  return (
    <div className="w-full px-4 py-5 border-b overflow-x-auto">
      <div className="flex gap-4">
        {Array.from({ length: 10 }).map((_, i) => (
          // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
          <Call key={i} />
        ))}
      </div>
    </div>
  );
}
