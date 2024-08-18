import Navigation from "../components/common/Navigation";
import PostList from "../components/posts/PostsList";

export default async function Page() {
    return (
        <main>
            <Navigation />
            <PostList />
        </main>
    )
}