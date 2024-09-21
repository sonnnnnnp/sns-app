
const Page = ({ params }: { params: { username: string }}) => {
    return <h1>{params.username}</h1>
};

export default Page;