import Layout from "@/layouts/layout";
import Messages from "@/app/messages/Messages";
import { HiddenWidget } from "@/components/common/hidden-widget";

const Page = () => {
  return (
    <Layout>
      <main className="relative flex-grow max-w-[750px]">
        <Messages />
      </main>
      <HiddenWidget />
    </Layout>
  );
};

export default Page;
