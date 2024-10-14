import Layout from "@/layouts/layout";
import Messages from "@/app/messages/Messages";
import { Widget } from "@/components/common/widget";

const Page = () => {
  return (
    <Layout>
      <main className="flex-grow max-w-[750px]">
        <Messages />
      </main>
      <Widget />
    </Layout>
  );
};

export default Page;
