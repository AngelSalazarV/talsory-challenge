export interface MatrixStats {
  max: number;
  min: number;
  average: number;
  total: number;
  isDiagonal: boolean;
}

export function calculateStats(matrix: number[][]): MatrixStats {
  const values = matrix.flat();

  const total = values.reduce((sum, value) => sum + value, 0);

  const max = Math.max(...values);
  const min = Math.min(...values);
  const average = total / values.length;

  const isDiagonal = matrix.every((row, i) =>
    row.every((value, j) => {
      if (i === j) {
        return true;
      }

      return value === 0;
    })
  );

  return {
    max,
    min,
    average,
    total,
    isDiagonal,
  };
}