import Layout from "@/layouts/layout";
import { Widget } from "@/components/common/widget";
import { Profile } from "./components/profile";

const ProfilePage = () => {
  return (
    <Layout>
      <main className="flex-grow max-w-[750px]">
        <Profile />
      </main>
      <Widget />
    </Layout>
  );
};

export default ProfilePage;
