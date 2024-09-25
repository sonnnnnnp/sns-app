import Navigation from "@/components/common/Navigation";
import { Timeline } from "./components/timeline";
import { Nav } from "@/components/common/nav";

const Page = () => {
  return (
    <div className="flex h-dvh justify-center">
      <Navigation />
      {/* <Nav /> */}
      <main className="flex justify-center w-[950px]">
        <Timeline />
      </main>
    </div>
  );
};
export default Page;
