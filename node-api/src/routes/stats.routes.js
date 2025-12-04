const express = require('express');
const router = express.Router();

const { postStats } = require('../controllers/stats.controller');
const jwtMiddleware = require('../middlewares/jwt.middleware');

// Protegemos /api/stats con JWT
router.post('/stats', jwtMiddleware, postStats);

module.exports = router;
