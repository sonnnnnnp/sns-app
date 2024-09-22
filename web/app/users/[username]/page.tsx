import { useFetchWithTokens } from "@/lib/useFetchWithToken";
const Page = ({ params }: { params: { username: string } }) => {
  useFetchWithTokens("/timeline");
  return (
    <>
      <h1>{params.username}</h1>
      <h2>{JSON.stringify(params)}</h2>
    </>
  );
};

export default Page;
