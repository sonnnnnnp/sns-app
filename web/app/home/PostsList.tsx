import Image from 'next/image';
import React from 'react';

interface Post {
    id: number;
    user_id: number;
    context: string;
    created_at: number;
    likes_count: number;
    comments_count: number;
}

// 改行コードを<br>タグに変換する関数
function convertNewlinesToBr(text: string) {
    return text.split('\n').map((line, index) => (
        <React.Fragment key={index}>
            {line}
            {index !== text.split('\n').length - 1 && <br />}
        </React.Fragment>
    ));
}

export default async function PostList() {
    const res = await fetch('http://localhost:3000/api/posts', {
        cache: 'no-store',
    });
    const posts: Post[] = await res.json();

    return (
        <div className="max-w-2xl">
            {posts.map( post => (
                <div key={post.id} className="flex p-3 bg-zinc-900 rounded-lg mb-2">
                    <div className="mr-2">
                        <a href={`user/${post.user_id}`} className="block w-12 h-12 relative">
                            <Image
                                className="rounded-full"
                                src="/kaworu_icon.jpg"
                                alt="profile icon"
                                fill
                            />
                        </a>
                    </div>
                    <div className="w-full">
                        <div className="flex items-center">
                            <a href={`user/${post.user_id}`} className="font-semibold">カヲル</a>
                            <span className="text-sm">@{post.user_id}</span>
                            <span className="material-symbols-outlined ml-auto">more_horiz</span>
                        </div>
                        <p className="">{convertNewlinesToBr(post.context)}</p>
                    </div>
                </div>
            ))}
        </div>
    )
}