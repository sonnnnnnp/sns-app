'use client';
type FeedToggleProps = {
    activeTab: 'recommendations' | 'following';
    setActiveTab: (tab: 'recommendations' | 'following') => void;
  };
  
  const FeedToggle: React.FC<FeedToggleProps> = ({ activeTab, setActiveTab }) => {
    return (
      <nav className="flex justify-center my-4">
        <button
          className={`px-4 py-2 mx-2 ${activeTab === 'recommendations' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setActiveTab('recommendations')}
        >
          おすすめ
        </button>
        <button
          className={`px-4 py-2 mx-2 ${activeTab === 'following' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setActiveTab('following')}
        >
          フォロー中
        </button>
      </nav>
    );
  };
  
  export default FeedToggle;