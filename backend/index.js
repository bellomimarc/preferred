"use strict";
const { build } = require("./app.js");

async function main() {
  // Run the server!
  try {
    const app = await build({
      logger: true,
    });

    await app.listen({ port: process.env.PORT });
  } catch (err) {
    console.error(err)
    process.exit(1);
  }
}

main();
