import Navigation from "@/components/common/Navigation";
import { Timeline } from "./components/timeline";
import { Nav } from "@/components/common/nav";

const Page = () => {
  return (
    <div className="flex h-dvh justify-center">
      <Navigation />
      {/* <Nav /> */}
      <main className="">
        <Timeline />
      </main>
    </div>
  );
};
export default Page;
