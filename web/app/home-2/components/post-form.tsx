"use client";
import { FormEvent } from "react";
import { Card } from "@/components/ui/card";
import client from "@/lib/api";

export function PostForm() {
  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const { data } = await client.POST("/posts/create", {
      body: {
        content: "rachmaninoff",
      },
    });
  }
  return (
    <Card>
      <form onSubmit={onSubmit}>
        <div></div>
        <textarea
          className="block p-3 w-full"
          //   value={content}
          //   onChange={(e) => setContent(e.target.value)}
          placeholder="投稿内容を入力する"
          rows={5}
          cols={40}
        />
        <div className="flex p-2">
          <button
            type="submit"
            className="bg-blue-500 text-white py-2 px-8 rounded-full ml-auto"
          >
            投稿する
          </button>
        </div>
      </form>
    </Card>
  );
}
