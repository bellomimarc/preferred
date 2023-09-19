const { expect, afterAll, beforeAll } = require('@jest/globals');
const { getConnectedClient } = require('./redis')

let client;

beforeAll(async () => {
    client = await getConnectedClient();
})

afterAll(async () => {
    await client.disconnect();
})

it('I can connect to redis', async () => {
    await client.set("foo", "1");

    const fooValue = await client.get("foo")

    expect(fooValue).toStrictEqual("1")
})

