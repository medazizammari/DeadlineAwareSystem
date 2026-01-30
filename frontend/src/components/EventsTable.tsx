import type { FC } from "react";
import type { RealTimeEvent } from "../types/types";

const formatDate = (date?: string) => {
  if (!date) return "-";
  const d = new Date(date);
  if (Number.isNaN(d.getTime())) return "-";
  return d.toLocaleString();
};

const EventsTable: FC<{ events: RealTimeEvent[] }> = ({ events }) => {
  return (
    <div style={{ overflowX: "auto", marginTop: 12 }}>
      <table style={{ width: "100%", borderCollapse: "collapse" }}>
        <thead>
          <tr style={{ textAlign: "left" }}>
            <th style={{ padding: 8, borderBottom: "1px solid #ccc" }}>ID</th>
            <th style={{ padding: 8, borderBottom: "1px solid #ccc" }}>Status</th>
            <th style={{ padding: 8, borderBottom: "1px solid #ccc" }}>Created</th>
            <th style={{ padding: 8, borderBottom: "1px solid #ccc" }}>Processed</th>
            <th style={{ padding: 8, borderBottom: "1px solid #ccc" }}>Deadline (ms)</th>
          </tr>
        </thead>
        <tbody>
          {events.map((e, i) => (
            <tr key={e.id + i}>
              <td style={{ padding: 8, borderBottom: "1px solid #eee" }}>
                {e.id.slice(0, 8)}
              </td>
              <td style={{ padding: 8, borderBottom: "1px solid #eee" }}>{e.status}</td>
              <td style={{ padding: 8, borderBottom: "1px solid #eee" }}>
                {formatDate(e.createdAt)}
              </td>
              <td style={{ padding: 8, borderBottom: "1px solid #eee" }}>
                {formatDate(e.processedAt)}
              </td>
              <td style={{ padding: 8, borderBottom: "1px solid #eee" }}>
                {Math.round(e.deadlineMs)}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default EventsTable;
