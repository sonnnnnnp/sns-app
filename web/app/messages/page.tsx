import Navigation from "@/components/common/Navigation";
import Messages from "@/app/messages/Messages";
const Page = () => {
  return (
    <div className="flex justify-center">
      <Navigation />
      <main>
        <Messages />
      </main>
    </div>
  );
};
export default Page;
