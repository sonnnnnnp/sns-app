import { Button } from "@nextui-org/react";
import { useNavigate } from "react-router";
import { SolarAltArrowLeftLinear, SolarMenuDotsBold } from "~/components/icons";
import UserContent from "~/components/users/user-content";
import UserProfile from "~/components/users/user-profile";

export default function User() {
  const navigate = useNavigate();

  return (
    <div className="w-full">
      <div className="sticky top-0 z-10 flex items-center px-2 w-full h-14 bg-transparent backdrop-blur-md border-b">
        <Button
          isIconOnly
          variant="light"
          radius="full"
          onPress={() => navigate(-1)}
        >
          <SolarAltArrowLeftLinear className="w-5 h-5" />
        </Button>
        <div className="grid gap-1 ml-2">
          <span className="font-bold">ユーザー名</span>
          <span className="text-xs text-foreground-400">214件の投稿</span>
        </div>
        <Button isIconOnly variant="light" radius="full" className="ml-auto">
          <SolarMenuDotsBold className="w-5 h-5 rotate-90" />
        </Button>
      </div>
      <UserProfile />
      <UserContent />
    </div>
  );
}
