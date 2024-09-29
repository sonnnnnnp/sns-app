const Page = ({ params }: { params: { username: string } }) => {
  return (
    <div>
      <h1>{params.username}</h1>
      <h2>{JSON.stringify(params)}</h2>
    </div>
  );
};

export default Page;
