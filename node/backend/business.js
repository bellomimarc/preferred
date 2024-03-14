const doSomething = (body) => {
  console.log(body);

  return { hello: "world post" };
};

const doMath = (body) => {
  const { a, b } = body;

  return { result: a + b };
};

module.exports = { doSomething, doMath };