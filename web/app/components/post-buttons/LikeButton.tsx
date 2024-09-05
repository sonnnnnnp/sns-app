'use client';
import React from 'react';

interface LikeButtonProps {
    likesCount: number;
}

const LikeButton: React.FC<LikeButtonProps> = ({ likesCount }) => (
    <div className="flex">
        <span className="material-symbols-outlined mr-1">thumb_up</span>
        <span className="w-8">{likesCount}</span>
    </div>
);

export default LikeButton;
