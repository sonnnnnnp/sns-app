import Navigation from "@/components/common/tmp-navigation";
import { Timeline } from "./components/timeline";

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
