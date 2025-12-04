const jwt = require('jsonwebtoken');
const { JWT_SECRET } = require('../config');
const { error } = require('../utils/response');

function jwtMiddleware(req, res, next) {
  // Si no hay secret, no exigimos token (modo dev)
  if (!JWT_SECRET) {
    return next();
  }

  const authHeader = req.headers['authorization'];
  if (!authHeader) {
    return res.status(401).json(
      error('missing Authorization header')
    );
  }

  const parts = authHeader.split(' ');
  if (parts.length !== 2 || parts[0].toLowerCase() !== 'bearer') {
    return res.status(401).json(
      error('invalid Authorization header format')
    );
  }

  const token = parts[1];

  try {
    const decoded = jwt.verify(token, JWT_SECRET);
    req.user = decoded; // claims
    return next();
  } catch (err) {
    return res.status(401).json(
      error('invalid or expired token')
    );
  }
}

module.exports = jwtMiddleware;
