import { calculateStatistics } from '../services/stats.service.js';
import { validateMatrices } from '../utils/matrix.utils.js';

async function getStatistics(req, res) {
  try {
    const validation = validateMatrices(req.body);
    
    if (!validation.valid) {
      return res.status(400).json({
        success: false,
        error: 'Datos de entrada inválidos',
        details: validation.errors
      });
    }

    const stats = calculateStatistics(req.body);

    return res.status(200).json({
      success: true,
      data: stats
    });

  } catch (error) {
    console.error('Error calculating statistics:', error);
    return res.status(500).json({
      success: false,
      error: 'Error interno del servidor',
      message: error.message
    });
  }
}

export { getStatistics };