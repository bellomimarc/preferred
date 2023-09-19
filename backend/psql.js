const { Pool } = require("pg");

const getConnectedClient = async () => {
  const pool = new Pool({
    host: "localhost",
    port: 5555,
    user: "livecoding",
    password: "livecoding",
    query_timeout: 3000,
    statement_timeout: 1500,
    max: 20,
    idleTimeoutMillis: 3000,
    connectionTimeoutMillis: 2000,
  });

  const client = await pool.connect();

  client.on('error', (err) => {
    console.error('something bad has happened!', err)
  })

  return client
};

module.exports = { getConnectedClient };
