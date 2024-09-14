type FeedToggleProps = {
  activeTab: "recommendations" | "following";
  setActiveTab: (tab: "recommendations" | "following") => void;
};

const FeedToggle: React.FC<FeedToggleProps> = ({ activeTab, setActiveTab }) => {
  return (
    <nav className="flex justify-center my-4 w-full max-w-2xl fixed z-10">
      <button
        className={`w-full px-4 py-2 mx-2 ${
          activeTab === "recommendations" ? "border-b-2" : ""
        }`}
        onClick={() => setActiveTab("recommendations")}
      >
        おすすめ
      </button>
      <button
        className={`w-full px-4 py-2 mx-2 ${
          activeTab === "following" ? "border-b-2" : ""
        }`}
        onClick={() => setActiveTab("following")}
      >
        フォロー中
      </button>
    </nav>
  );
};

export default FeedToggle;
