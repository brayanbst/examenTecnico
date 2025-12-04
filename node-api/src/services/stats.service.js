// Helper para errores con statusCode
function createError(statusCode, message) {
  const err = new Error(message);
  err.statusCode = statusCode;
  return err;
}

function isDiagonal(matrix, epsilon = 1e-9) {
  const rows = matrix.length;
  if (rows === 0) return false;
  const cols = matrix[0].length;
  if (rows !== cols) return false;

  for (let i = 0; i < rows; i++) {
    if (matrix[i].length !== cols) return false;
    for (let j = 0; j < cols; j++) {
      if (i !== j && Math.abs(matrix[i][j]) > epsilon) {
        return false;
      }
    }
  }
  return true;
}

function validateMatrices(matrices) {
  if (!Array.isArray(matrices) || matrices.length === 0) {
    throw createError(400, 'matrices must be a non-empty array');
  }

  for (const mat of matrices) {
    if (!Array.isArray(mat) || mat.length === 0) {
      throw createError(400, 'each matrix must be a non-empty 2D array');
    }
    const cols = mat[0].length;
    for (const row of mat) {
      if (!Array.isArray(row) || row.length !== cols) {
        throw createError(400, 'each matrix must be rectangular');
      }
      for (const val of row) {
        if (typeof val !== 'number') {
          throw createError(400, 'all matrix values must be numbers');
        }
      }
    }
  }
}

function computeStats(matrices) {
  validateMatrices(matrices);

  let maxValue = Number.NEGATIVE_INFINITY;
  let minValue = Number.POSITIVE_INFINITY;
  let totalSum = 0;
  let count = 0;

  for (const mat of matrices) {
    for (const row of mat) {
      for (const val of row) {
        if (val > maxValue) maxValue = val;
        if (val < minValue) minValue = val;
        totalSum += val;
        count++;
      }
    }
  }

  const average = count > 0 ? totalSum / count : 0;
  const diagonals = matrices.map(m => isDiagonal(m));

  return {
    maxValue,
    minValue,
    average,
    totalSum,
    diagonals,
  };
}

module.exports = {
  computeStats,
};
