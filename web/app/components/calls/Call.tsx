import { Avatar, Button } from "@nextui-org/react";
import { SolarMicrophone2Linear } from "../icons/solar";

export function Call() {
  return (
    <div className="flex flex-col items-center flex-shrink-0 text-sm">
      <Button isIconOnly variant="light" className="w-20 h-20 mb-2">
        <div className="grid grid-cols-2 w-full h-full rounded-large overflow-hidden">
          {Array.from({ length: 4 }).map((_, i) => (
            <Avatar
              // biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
              key={i}
              radius="none"
              src="https://i.pravatar.cc/150?u=a042581f4e29026024d"
            />
          ))}
        </div>
      </Button>
      <span className="flex items-center gap-1 text-nowrap text-xs">
        <span className="p-0.5 bg-default rounded-full">
          <SolarMicrophone2Linear width={14} height={14} />
        </span>
        4人参加中
      </span>
    </div>
  );
}
