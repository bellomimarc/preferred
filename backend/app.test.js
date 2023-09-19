"use strict";

const { build } = require("./app.js");

jest.mock("./business.js", () => {
  return {
    doMath: jest.fn(() => {
      return { result: 30 };
    }),
    doSomething: jest.fn(() => {
      return { hello: "mock" };
    }),
  };
});

afterAll(() => {
  jest.restoreAllMocks();
});

it('requests the "/" route', async () => {
  const app = build();

  const response = await app.inject({
    method: "GET",
    url: "/",
  });

  expect(response.statusCode).toEqual(200);
  expect(response.json()).toEqual({ hello: "world" });
});

it('requests the "/" route with a POST', async () => {
  const app = build();

  const response = await app.inject({
    method: "POST",
    url: "/",
  });

  expect(response.statusCode).toEqual(200);
  expect(response.json()).toEqual({ hello: "mock" });
});

it('requests the "/math" route with a POST', async () => {
  const app = build();

  const response = await app.inject({
    method: "POST",
    url: "/math",
    payload: {
      a: 1,
      b: 2,
    },
  });

  expect(response.statusCode).toEqual(200);
  expect(response.json()).toEqual({ result: 30 });
});
