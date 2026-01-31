import useRealtimeEvents from "./hooks/useRealtimeEvents";
import StatsBar from "./components/StatsBar";
import EventsTable from "./components/EventsTable";

export default function App() {
  const events = useRealtimeEvents();

  async function createEvent() {
    try {
      const res = await fetch(`${import.meta.env.VITE_API_HTTP}/event`, { method: "POST" });
      if (!res.ok) {
        const text = await res.text();
        alert(`Create event failed: ${res.status} ${text}`);
      }
    } catch {
      alert("Backend not reachable (is Go server running on :8080?)");
    }
  }

  return (
    <div style={{ maxWidth: 1000, margin: "0 auto", padding: 16 }}>
      <h1 style={{ marginBottom: 12 }}>Real-Time Event Monitor</h1>

      <button onClick={createEvent} style={{ marginBottom: 12 }}>
        Send event
      </button>

      <StatsBar events={events} />
      <EventsTable events={events} />
    </div>
  );
}
