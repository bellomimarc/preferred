const path = require("path");
const { doMath, doSomething } = require("./business");
const { getConnectedClient } = require("./redis");

const build = async (opts) => {
  const fastify = require("fastify")(opts);

  fastify.register(require("@fastify/static"), {
    root: path.join(process.cwd(), "public"),
    prefix: "/public/",
  });

  const redisClient = await getConnectedClient()
  fastify.decorate("redis", redisClient);
  fastify.addHook("onClose", () => {
    return redisClient.disconnect()
  })

  fastify.get("/", async function handler(request, reply) {
    return { hello: "world" };
  });

  fastify.post("/", async function handler(request, reply) {
    const response = doSomething(request.body);

    reply.send(response);
  });

  const schema = {
    body: {
      type: "object",
      required: ["a", "b"],
      properties: {
        a: { type: "number" },
        b: { type: "number" },
      },
    },
    response: {
      200: {
        type: "object",
        properties: {
          result: { type: "number" },
        },
      },
    },
  };
  fastify.post("/math", { schema }, async function handler(request, reply) {
    const { a, b } = request.body;
    const { result } = doMath({ a, b });

    reply.send({ result });
  });

  return fastify;
};

module.exports = { build };
