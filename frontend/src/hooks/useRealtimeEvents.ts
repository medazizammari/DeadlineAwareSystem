import { useEffect, useRef, useState } from "react";
import type { RealTimeEvent } from "../types/types";

export default function useRealtimeEvents() {
  const [events, setEvents] = useState<RealTimeEvent[]>([]);
  const buffer = useRef<RealTimeEvent[]>([]);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onmessage = (msg) => {
      try {
        buffer.current.push(JSON.parse(msg.data));
      } catch {
        // ignore malformed message
      }
    };

    let rafId = 0;

    const flush = () => {
      if (buffer.current.length > 0) {
        const pending = buffer.current.splice(0, buffer.current.length);

        setEvents((prev) => {
          const next = [...prev, ...pending];
          return next.slice(-50); // keep last 50 to avoid UI slowdown
        });
      }

      rafId = requestAnimationFrame(flush);
    };

    rafId = requestAnimationFrame(flush);

    return () => {
      ws.close();
      cancelAnimationFrame(rafId);
    };
  }, []);

  return events;
}
