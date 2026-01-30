import type { FC } from "react";
import type { RealTimeEvent } from "../types/types";

const StatsBar: FC<{ events: RealTimeEvent[] }> = ({ events }) => {
  const late = events.filter((e) => e.status === "late").length;
  const dropped = events.filter((e) => e.status === "dropped").length;

  return (
    <div style={{ padding: "10px", background: "#eee" }}>
      <strong>Events: {events.length}</strong>{" "}
      <strong style={{ marginLeft: 12 }}>Late: {late}</strong>{" "}
      <strong style={{ marginLeft: 12 }}>Dropped: {dropped}</strong>
    </div>
  );
};

export default StatsBar;
