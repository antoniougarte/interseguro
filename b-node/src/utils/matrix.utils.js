function getAllValues(matrices) {
  const values = [];
  
  for (const matrix of Object.values(matrices)) {
    if (!Array.isArray(matrix)) continue;
    
    for (const row of matrix) {
      if (!Array.isArray(row)) continue;
      values.push(...row.filter(v => typeof v === 'number' && !isNaN(v)));
    }
  }
  
  return values;
}

function getMax(values) {
  if (values.length === 0) return null;
  return Math.max(...values);
}

function getMin(values) {
  if (values.length === 0) return null;
  return Math.min(...values);
}

function getAverage(values) {
  if (values.length === 0) return null;
  const sum = values.reduce((acc, val) => acc + val, 0);
  return sum / values.length;
}

function getSum(values) {
  if (values.length === 0) return null;
  return values.reduce((acc, val) => acc + val, 0);
}

function isDiagonal(matrix) {
  if (!Array.isArray(matrix) || matrix.length === 0) {
    return false;
  }

  const rows = matrix.length;
  const cols = matrix[0]?.length || 0;

  if (rows !== cols) {
    return false;
  }

  const EPSILON = 1e-10;
  
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (i !== j) {
        if (Math.abs(matrix[i][j]) > EPSILON) {
          return false;
        }
      }
    }
  }

  return true;
}

function validateMatrices(data) {
  const errors = [];

  if (!data || typeof data !== 'object') {
    errors.push('El cuerpo de la solicitud debe ser un objeto');
    return { valid: false, errors };
  }

  const { rotated, Q, R } = data;

  if (!rotated && !Q && !R) {
    errors.push('Debe proporcionar al menos una matriz (rotated, Q, o R)');
  }

  const matrices = { rotated, Q, R };
  
  for (const [name, matrix] of Object.entries(matrices)) {
    if (matrix !== undefined && matrix !== null) {
      if (!Array.isArray(matrix)) {
        errors.push(`${name} debe ser un array`);
        continue;
      }

      if (matrix.length === 0) {
        errors.push(`${name} no puede estar vacío`);
        continue;
      }

      const firstRowLength = matrix[0]?.length;
      
      for (let i = 0; i < matrix.length; i++) {
        if (!Array.isArray(matrix[i])) {
          errors.push(`${name}[${i}] debe ser un array`);
          break;
        }

        if (matrix[i].length !== firstRowLength) {
          errors.push(`${name} debe ser rectangular (todas las filas deben tener el mismo tamaño)`);
          break;
        }

        for (let j = 0; j < matrix[i].length; j++) {
          if (typeof matrix[i][j] !== 'number' || isNaN(matrix[i][j])) {
            errors.push(`${name}[${i}][${j}] debe ser un número válido`);
            break;
          }
        }
      }
    }
  }

  return {
    valid: errors.length === 0,
    errors
  };
}

export {
  getAllValues,
  getMax,
  getMin,
  getAverage,
  getSum,
  isDiagonal,
  validateMatrices
};