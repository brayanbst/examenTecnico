const { success } = require('../utils/response');
const { computeStats } = require('../services/stats.service');

// Controlador para POST /api/stats
async function postStats(req, res, next) {
  try {
    const { matrices } = req.body;

    const stats = computeStats(matrices);

    return res.status(200).json(
      success('statistics computed successfully', stats)
    );
  } catch (err) {
    // Delegamos al middleware de errores
    return next(err);
  }
}

module.exports = {
  postStats,
};
