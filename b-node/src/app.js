import express from 'express';
import cors from 'cors';
import { getStatistics } from './controllers/stats.controller.js';
import { authMiddleware } from './middlewares/auth.middleware.js';

const app = express();

app.use(cors());
app.use(express.json());

app.use((req, res, next) => {
  console.log(`[${new Date().toISOString()}] ${req.method} ${req.path}`);
  next();
});

// Rutas públicas
app.get('/api/health', (req, res) => {
  res.json({
    status: 'ok',
    message: 'Node.js API funcionando correctamente',
    timestamp: new Date().toISOString()
  });
});

// Rutas protegidas
app.post('/api/stats', authMiddleware, getStatistics);

app.use((req, res) => {
  res.status(404).json({
    success: false,
    error: 'Ruta no encontrada'
  });
});

app.use((err, req, res, next) => {
  console.error('Error:', err);
  res.status(500).json({
    success: false,
    error: 'Error interno del servidor',
    message: err.message
  });
});

export default app;