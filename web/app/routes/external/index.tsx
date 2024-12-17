import { Button, Link } from "@nextui-org/react";

export function meta() {
  return [
    { title: "Reverie" },
    { name: "description", content: "Welcome to Reverie!" },
  ];
}

export default function Index() {
  return (
    <main className="grid place-items-center h-dvh w-full">
      <Button as={Link} href="/login" color="primary">
        ログイン
      </Button>
    </main>
  );
}
