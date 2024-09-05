'use client';
import React, { useState, useCallback } from 'react';

interface LikeButtonProps {
    likesCount: number;
    postId: number;
    onLike: (postId: number, isLiked: boolean) => Promise<void>;
}

const LikeButton: React.FC<LikeButtonProps> = ({ likesCount, postId, onLike }) => {
    const [isLiked, setIsLiked] = useState(false);
    const [localLikesCount, setLocalLikesCount] = useState(likesCount);

    const handleClick = useCallback(async () => {
        try {
            await onLike(postId, isLiked);
            setIsLiked(!isLiked);
            setLocalLikesCount(prev => isLiked ? prev - 1 : prev + 1);
        } catch (error) {
            console.error('Error handling like:', error);
        }
    }, [isLiked, onLike, postId]);

    return (
        <div className="flex items-center cursor-pointer" onClick={handleClick}>
            <span className={`material-symbols-outlined mr-1 ${isLiked ? 'text-pink-500' : ''}`}>
                {isLiked ? 'favorite' : 'favorite_border'}
            </span>
            <span className="w-8">{localLikesCount}</span>
        </div>
    );
};

export default LikeButton;
