import { calculateStats } from "./stats.service";

describe("calculateStats", () => {
  test("calculates matrix statistics", () => {
    const matrix = [
      [1, 2],
      [3, 4],
    ];

    const result = calculateStats(matrix);

    expect(result.max).toBe(4);
    expect(result.min).toBe(1);
    expect(result.average).toBe(2.5);
    expect(result.total).toBe(10);
    expect(result.isDiagonal).toBe(false);
  });

  test("detects a diagonal matrix", () => {
    const matrix = [
      [1, 0],
      [0, 2],
    ];

    const result = calculateStats(matrix);

    expect(result.isDiagonal).toBe(true);
  });
});