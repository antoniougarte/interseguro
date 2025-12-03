import {
  getAllValues,
  getMax,
  getMin,
  getAverage,
  getSum,
  isDiagonal
} from '../utils/matrix.utils.js';

function calculateStatistics(data) {
  const { rotated, Q, R } = data;
  const allValues = getAllValues({ rotated, Q, R });

  const max = getMax(allValues);
  const min = getMin(allValues);
  const average = getAverage(allValues);
  const sum = getSum(allValues);

  const isDiagonalCheck = {
    rotated: rotated ? isDiagonal(rotated) : null,
    Q: Q ? isDiagonal(Q) : null,
    R: R ? isDiagonal(R) : null
  };

  return {
    max,
    min,
    average,
    sum,
    isDiagonal: isDiagonalCheck
  };
}

export { calculateStatistics };