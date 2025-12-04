const express = require('express');
const cors = require('cors');

const statsRoutes = require('./routes/stats.routes');
const errorMiddleware = require('./middlewares/error.middleware');

const app = express();

// Middlewares globales
app.use(cors());
app.use(express.json());

// Rutas bajo /api
app.use('/api', statsRoutes);

// Middleware de errores (al final)
app.use(errorMiddleware);

module.exports = app;
