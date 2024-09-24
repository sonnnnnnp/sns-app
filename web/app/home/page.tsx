import Navigation from "@/components/common/Navigation";
import { Timeline } from "./components/timeline";
import { Separator } from "@/components/ui/separator";

const Page = () => {
  return (
    <main className="flex h-dvh justify-center">
      <Navigation />
      <div>
        <Timeline />
      </div>
    </main>
  );
};

export default Page;
