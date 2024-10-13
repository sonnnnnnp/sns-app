import Layout from "@/layouts/layout";
import { Widget } from "@/components/common/widget";
import { Timeline } from "./components/timeline";

const HomePage = () => {
  return (
    <Layout>
      <main className="flex-grow max-w-[750px]">
        <Timeline />
      </main>
      <Widget />
    </Layout>
  );
};

export default HomePage;
