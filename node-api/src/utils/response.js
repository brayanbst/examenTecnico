// Respuesta de éxito: code = "000"
function success(message, data) {
  return {
    code: '000',
    message,
    data,
  };
}

// Respuesta de error: code = "001"
function error(message, data = null) {
  return {
    code: '001',
    message,
    data,
  };
}

module.exports = {
  success,
  error,
};
