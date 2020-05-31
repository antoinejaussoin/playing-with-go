console.log('Hello world');
console.time('exec')
console.log(calculatePi(1000000000))
console.timeEnd('exec')


function calculatePi(numberOfIterations) {
  // Pi/4 = 1 - 1/3 + 1/5 - 1/7 + ...
  let pi = 1;
  for (i = 1; i <= numberOfIterations; i++) {
    const isNegative = i % 2 == 0 ? false : true;
    const factor = i * 2 + 1
    if (isNegative) {
      pi -= 1 / factor
    } else {
      pi += 1 / factor
    }
  }
  return pi * 4
}