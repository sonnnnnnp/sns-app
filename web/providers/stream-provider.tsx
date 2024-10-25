"use client";

import { createContext, useContext } from "react";

type Stream = {
  connection: WebSocket | null;
};

const defaultStream: Stream = {
  connection: new WebSocket("ws://localhost:1323/stream"),
};

const StreamContext = createContext<Stream>(defaultStream);

export const StreamProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <StreamContext.Provider value={defaultStream}>
      {children}
    </StreamContext.Provider>
  );
};

export const useStream = () => useContext(StreamContext);
