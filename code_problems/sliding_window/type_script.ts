type Args = {
  correctLength: number;
  numbers: number[];
  targetValue: number;
};

/**
 * Find the largest list that is less than the target value.
 */
const slidingWindow = (args: Args): number => {
  const { numbers, targetValue } = args;
  const endPoint = numbers.length;
  let left = 0; //5
  let right = 0; //7
  let currentMaxSize = 0; //5

  while (right < endPoint) {
    console.log("start", currentMaxSize, left, right);
    // both pointers are in the same place
    if (left === right) {
      // all other numbers are greater
      if (sumArray(numbers.slice(left, right + 1)) < targetValue) {
        currentMaxSize = numbers.slice(left, right + 1).length;
      }
    } else {
      // less than target case
      if (
        sumArray(numbers.slice(left, right + 1)) < targetValue &&
        numbers.slice(left, right + 1).length > currentMaxSize
      ) {
        currentMaxSize = numbers.slice(left, right + 1).length;
        // pointers are side by side
      } else if (left !== right) {
        // move left pointer forward until it's less than
        while (
          left !== right &&
          sumArray(numbers.slice(left, right + 1)) > targetValue
        ) {
          left++;
        }
        const currentLength = numbers.slice(left, right + 1).length;
        if (
          sumArray(numbers.slice(left, right + 1)) < targetValue &&
          currentLength > currentMaxSize
        ) {
          currentMaxSize = currentLength;
        }
      }
    }
    console.log("end", currentMaxSize, left, right);
    right++;
  }
  return currentMaxSize;
};

const sumArray = (arr: number[]): number =>
  arr.reduce((currentSum, a) => currentSum + a, 0);

const testA: Args = {
  correctLength: 2,
  numbers: [1, 2, 3, 4],
  targetValue: 4,
};
const testB: Args = {
  correctLength: 2,
  numbers: [1, 2, 3, 4],
  targetValue: 6,
};
const testC: Args = {
  correctLength: 5,
  numbers: [10, 1, 3, 2, 4, 1, 8, 2],
  targetValue: 12,
};
const testD: Args = {
  correctLength: 1,
  numbers: [10, 1, 3, 2, 4, 1, 8, 2],
  targetValue: 2,
};

const runTest = (testName: string, args: Args) => {
  const result = slidingWindow(args);
  console.log(
    `Did ${testName} work: ${result === args.correctLength}, result ${result}`,
  );
};

runTest("testA", testA);
runTest("testB", testB);
runTest("testC", testC);
runTest("testD", testD);
