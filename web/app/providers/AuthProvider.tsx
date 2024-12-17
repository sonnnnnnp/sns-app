import Cookies from "js-cookie";
import React from "react";

const AuthContext = React.createContext(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = React.useState(null);

  React.useEffect(() => {
    console.log("check cookies");

    if (!Cookies.get("access-token")) {
      window.location.href = "/";
      return;
    }

    console.log("welcome user!");

    setUser(null);
  }, []);

  return <AuthContext.Provider value={user}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  React.useContext(AuthContext);
}
