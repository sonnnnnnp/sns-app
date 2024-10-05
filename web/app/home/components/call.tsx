import { CallData } from "./call-list";
import { Avatar, AvatarImage } from "@/components/ui/avatar";

export function Call({ call }: { call: CallData }) {
  return (
    <div className="flex flex-col items-center cursor-pointer">
      <div className="grid grid-cols-2 w-20 h-20 mb-2 rounded-lg overflow-hidden">
        {call.participants.slice(0, 4).map((participant, i) => (
          <Avatar key={i} className="rounded-none">
            <AvatarImage
              src={participant.avatar_url ?? "/users/placeholder-profile.svg"}
            />
          </Avatar>
        ))}
      </div>
      <span>{`${call.participants.length}人参加中`}</span>
    </div>
  );
}
