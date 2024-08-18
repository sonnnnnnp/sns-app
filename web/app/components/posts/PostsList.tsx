
interface Post {
    id: number;
    user_id: number;
    context: string;
    created_at: number;
    likes_count: number;
    comments_count: number;
}

export default async function PostList() {
    const res = await fetch('http://localhost:3001/api/posts', {
        cache: 'no-store',
    });
    const posts: Post[] = await res.json();

    return (
        <>
            {posts.map( post => (
                <p key={post.id} className="text-red-600">{post.context}</p>
            ))}
        </>
    )
}