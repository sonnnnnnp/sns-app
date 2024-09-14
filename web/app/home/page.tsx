import Navigation from "@/components/common/Navigation";
import PostList from "@/app/home/PostsList";

const Page = () => {
  return (
    <div className="flex justify-center">
      <Navigation />
      <main>
        <PostList />
      </main>
    </div>
  );
};
export default Page;
