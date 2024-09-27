import Navigation from "@/components/common/Navigation";
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
