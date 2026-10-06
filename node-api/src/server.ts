import express from "express";
import { getStats } from "./controllers/stats.controller";

const app = express();

app.use(express.json());

app.get("/health", (_req, res) => {
  res.json({
    status: "ok",
  });
});

app.post("/api/stats", getStats);

app.listen(3001, () => {
  console.log("Node API running on port 3001");
});