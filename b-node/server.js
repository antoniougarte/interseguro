import dotenv from 'dotenv';
import app from './src/app.js';

dotenv.config();

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`✅ Servidor Node.js ejecutándose en http://localhost:${PORT}`);
  console.log(`📊 Endpoint de estadísticas: POST http://localhost:${PORT}/api/stats`);
  console.log(`💚 Health check: GET http://localhost:${PORT}/api/health`);
});