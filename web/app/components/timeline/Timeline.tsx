import { CallTimeline } from "./CallTimeline";
import { PostTimeline } from "./PostTimeline";

export function Timeline({ type }: { type: "following" | "public" }) {
  return (
    <div>
      <CallTimeline />
      <PostTimeline type={type} />
    </div>
  );
}
