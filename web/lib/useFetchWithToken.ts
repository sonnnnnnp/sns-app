import { cookies } from "next/headers";
import { client } from "./api";

export const useFetchWithTokens = async (path: string) => {
  const cookiesStore = cookies();
  const accesToken = cookiesStore.get("access-token")!.value;
  const response = await client.GET(path, {
    headers: {
      // Authorization: `Bearer ${accesToken}`,
    },
  });
  console.log(response);
  return response;
};
