import React from "react";

interface RepostButtonProps {
  repostCount: number;
}

const RepostButton: React.FC<RepostButtonProps> = ({ repostCount }) => (
  <div className="flex">
    <span className="material-symbols-outlined mr-1">refresh</span>
    <span className="w-8">{repostCount}</span>
  </div>
);

export default RepostButton;
