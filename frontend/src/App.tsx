import useRealtimeEvents from "./hooks/useRealtimeEvents";
import StatsBar from "./components/StatsBar";
import EventsTable from "./components/EventsTable";

export default function App() {
  const events = useRealtimeEvents();

  return (
    <div style={{ maxWidth: 1000, margin: "0 auto", padding: 16 }}>
      <h1 style={{ marginBottom: 12 }}>Real-Time Event Monitor</h1>
      <StatsBar events={events} />
      <EventsTable events={events} />
    </div>
  );
}