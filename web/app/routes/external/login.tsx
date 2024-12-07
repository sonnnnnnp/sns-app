import Cookies from "js-cookie";
import React from "react";

export default function Login() {
  React.useEffect(() => {
    if (Cookies.get("access-token")) {
      window.location.href = "/home";
      return;
    }
  }, []);

  return (
    <div className="flex w-full items-center justify-center">ログイン</div>
  );
}
