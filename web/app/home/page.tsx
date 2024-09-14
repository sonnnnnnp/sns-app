import Navigation from "@/components/common/Navigation";
import { Timeline } from "./components/timeline";

const Page = () => {
  return (
    <div className="flex h-dvh justify-center">
      <Navigation />
      <main>
        <Timeline />
      </main>
    </div>
  );
};
export default Page;
