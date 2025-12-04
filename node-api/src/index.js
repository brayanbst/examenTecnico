const { PORT } = require('./config');
const app = require('./app');

app.listen(PORT, () => {
  console.log(`Node API listening on port ${PORT}`);
});