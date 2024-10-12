import Layout from "@/layouts/layout";
import Timeline from "./components/timeline";
import FooterNavigation from "@/components/common/footer-navigation";

const HomePage = () => {
  return (
    <Layout>
      <main className="flex-grow max-w-[750px]">
        <Timeline />
        <FooterNavigation />
      </main>
    </Layout>
  );
};

export default HomePage;
