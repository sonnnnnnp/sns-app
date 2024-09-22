import { cookies } from "next/headers";

export const useFetchWithTokens = async (path: string) => {
  const cookiesStore = cookies();
  const accesToken = cookiesStore.get("access-token")!.value;
  const response = await fetch(path, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${accesToken}`,
      "Content-Type": "application/json",
    },
  });
  const data = await response.json();

  return data;
};
