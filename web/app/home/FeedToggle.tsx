type FeedToggleProps = {
  activeTab: "recommendations" | "following";
  setActiveTab: (tab: "recommendations" | "following") => void;
};

const FeedToggle: React.FC<FeedToggleProps> = ({ activeTab, setActiveTab }) => {
  return (
    <div className="bg-white bg-opacity-50 backdrop-blur-sm fixed z-10 w-full max-w-2xl">
      <nav className="flex justify-center my-4">
        <button
          className={`w-full px-4 py-2 mx-2 ${
            activeTab === "recommendations"
              ? "border-b-2 border-black/50 font-bold"
              : ""
          }`}
          onClick={() => setActiveTab("recommendations")}
        >
          おすすめ
        </button>
        <button
          className={`w-full px-4 py-2 mx-2 ${
            activeTab === "following"
              ? "border-b-2 border-black/50 font-bold"
              : ""
          }`}
          onClick={() => setActiveTab("following")}
        >
          フォロー中
        </button>
      </nav>
    </div>
  );
};

export default FeedToggle;
