const { error } = require('../utils/response');

function errorMiddleware(err, req, res, next) {
  console.error('Unexpected error:', err);

  const status = err.statusCode || 500;
  const message = err.message || 'internal server error';

  return res.status(status).json(
    error(message)
  );
}

module.exports = errorMiddleware;
