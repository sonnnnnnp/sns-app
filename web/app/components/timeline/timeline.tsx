import { CallTimeline } from "./call-timeline";
import { PostTimeline } from "./post-timeline";

export function Timeline() {
  return (
    <div>
      <CallTimeline />
      <PostTimeline />
    </div>
  );
}
