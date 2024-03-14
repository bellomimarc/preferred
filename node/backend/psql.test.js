const { beforeAll, afterAll, expect } = require("@jest/globals");
const { getConnectedClient } = require("./psql");

let client;

beforeAll(async () => {
    client = await getConnectedClient();
})

afterAll(async () => {
    await client.end();
})

it('I can execute a SELECT 1; statement', async () => {
    const { rowCount, rows } = await client.query('SELECT 1 as r;')

    expect(rowCount).toEqual(1);
    expect(rows[0].r).toEqual(1);
})
