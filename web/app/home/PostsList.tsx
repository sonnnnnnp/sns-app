import Image from 'next/image';
import React, { useState } from 'react';
import { CommentButton, RepostButton, LikeButton } from '../components/post-buttons/index';

interface Post {
    id: number;
    user_id: number;
    context: string;
    created_at: number;
    comments_count: number;
    repost_count: number;
    likes_count: number;
}

// 改行コードを<br>タグに変換する関数
function convertNewlinesToBr(text: string) {
    const lines = text.split('\n');
    return lines.map((line, index) => (
                <React.Fragment key={index}>
                    {line}
                    {index < lines.length - 1 && <br />}
                </React.Fragment>
            ))
}

export default async function PostList() {
    const res = await fetch(`${process.env.API_URL}/posts`, {
        cache: 'no-store',
    });
    const posts: Post[] = await res.json();
    const [activeTab, setActiveTab] = useState<'recommendations'|'following'>('recommendations');
    return (
        <div className="max-w-2xl">
            <nav>
                <button></button>
            </nav>
            {posts.map( post => (
                <div key={post.id} className="flex p-3 bg-stone-50 dark:bg-zinc-900 rounded-lg mb-2">
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
                        
                        <p className="mr-4">{convertNewlinesToBr(post.context)}</p>
                        
                        <div className="flex justify-between max-w-44">
                            <CommentButton commentsCount={post.comments_count} />
                            <RepostButton repostCount={post.repost_count} />
                            <LikeButton likesCount={post.likes_count} />
                        </div>
                    </div>
                </div>
            ))}
        </div>
    )
}