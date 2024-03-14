const fp = require("fastify-plugin");
const { getConnectedClient } = require("./redis");

module.exports = fp(
  (fastify, opts, done) => {
    getConnectedClient().then((client) => {
      fastify.decorate("redis", client);
      fastify.addHook("onClose", () => {
        return client.disconnect();
      });

      done();
    });
  },
  { fastify: "4.x", name: "redis" }
);
