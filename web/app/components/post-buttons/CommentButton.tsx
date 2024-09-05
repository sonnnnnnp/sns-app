import React from 'react';

interface CommentButtonProps {
    commentsCount: number;
}

const CommentButton: React.FC<CommentButtonProps> = ({ commentsCount }) => (
    <div className="flex">
        <span className="material-symbols-outlined mr-1">notes</span>
        <span className="w-8">{commentsCount}</span>
    </div>
);

export default CommentButton;
