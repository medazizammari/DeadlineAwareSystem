import { useEffect, useRef, useState } from "react";
import type { RealTimeEvent } from "../types/types";

// 🔒 WebSocket singleton (one per browser tab)
let ws: WebSocket | null = null;

function getWebSocket() {
  if (
    ws &&
    (ws.readyState === WebSocket.OPEN ||
      ws.readyState === WebSocket.CONNECTING)
  ) {
    return ws;
  }

  ws = new WebSocket("ws://localhost:8080/ws");
  return ws;
}

export default function useRealtimeEvents() {
  const [events, setEvents] = useState<RealTimeEvent[]>([]);
  const buffer = useRef<RealTimeEvent[]>([]);

  useEffect(() => {
    const socket = getWebSocket();

    const onMessage = (msg: MessageEvent) => {
      try {
        buffer.current.push(JSON.parse(msg.data));
      } catch {
        // ignore malformed message
      }
    };

    socket.addEventListener("message", onMessage);

    let rafId = 0;

    const flush = () => {
      if (buffer.current.length > 0) {
        const pending = buffer.current.splice(0, buffer.current.length);

        setEvents((prev) => {
          const next = [...prev, ...pending];
          return next.slice(-50); // keep last 50
        });
      }

      rafId = requestAnimationFrame(flush);
    };

    rafId = requestAnimationFrame(flush);

    return () => {
      socket.removeEventListener("message", onMessage);
      cancelAnimationFrame(rafId);
      // ❗ DO NOT close the WebSocket here
      // closing here causes double consumers on refresh/dev reload
    };
  }, []);

  return events;
}
