'use client';
import Navigation from '../components/common/Navigation';
import PostList from './PostsList';

export default async function Page() {
    return (
        <div className="flex justify-center">
            <Navigation />
            <main>
                <PostList />
            </main>
        </div>
    )
}