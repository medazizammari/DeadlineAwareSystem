export type RealTimeEvent = {
  id: string;
  createdAt: string;
  processedAt?: string;
  deadlineMs: number;
  status: "on-time" | "late" | "dropped";
};