import { Card } from "@/components/ui/card";
export function PostForm() {
  return (
    <Card>
      <h1>新規投稿</h1>
      {/* <form onSubmit={handleSubmit}> */}
      <textarea
        //   value={content}
        //   onChange={(e) => setContent(e.target.value)}
        placeholder="投稿内容を入力してください"
        rows={5}
        cols={40}
      />
      <button type="submit">送信</button>
      {/* </form> */}
    </Card>
  );
}
