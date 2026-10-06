import { Request, Response } from "express";
import { calculateStats } from "../services/stats.service";

export function getStats(req: Request, res: Response) {
  const { matrix } = req.body;

  if (!Array.isArray(matrix)) {
    return res.status(400).json({
      error: "matrix must be an array",
    });
  }

  const result = calculateStats(matrix);

  return res.json(result);
}