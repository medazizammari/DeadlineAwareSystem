import { useEffect, useRef, useState } from "react";
import type { RealTimeEvent } from "../types/types";

// One WS per browser tab
let ws: WebSocket | null = null;

function buildWsUrl() {
  const scheme = window.location.protocol === "https:" ? "wss" : "ws";
  return `${scheme}://${window.location.host}/ws`;
}

function getWebSocket() {
  if (
    ws &&
    (ws.readyState === WebSocket.OPEN ||
      ws.readyState === WebSocket.CONNECTING)
  ) {
    return ws;
  }

  ws = new WebSocket(buildWsUrl());

  ws.addEventListener("close", () => {
    ws = null;
  });
  ws.addEventListener("error", () => {
    ws = null;
  });

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
        setEvents((prev) => [...prev, ...pending].slice(-50));
      }
      rafId = requestAnimationFrame(flush);
    };

    rafId = requestAnimationFrame(flush);

    return () => {
      socket.removeEventListener("message", onMessage);
      cancelAnimationFrame(rafId);
    };
  }, []);

  return events;
}
