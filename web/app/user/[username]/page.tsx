const Page = ({ params }: { params: { username: string } }) => {
  return (
    <>
      <h1>{params.username}</h1>
      <h2>{JSON.stringify(params)}</h2>
    </>
  );
};

export default Page;
