import Navigation from "@/components/common/Navigation";
import Messages from "@/app/messages/Messages";
const Page = () => {
  return (
    <div className="flex justify-center">
      <Navigation />
      <main className="flex justify-center w-[950px]">
        <Messages />
      </main>
    </div>
  );
};
export default Page;
