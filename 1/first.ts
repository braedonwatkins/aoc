import * as fs from "fs";

const filePath = "./input.txt";

const getNumberFromLine = (line: string) => {
  const digits = line.match(/\d/g) || [];
  if (digits.length === 0) return 0;

  const firstDigit = digits[0];
  const lastDigit = digits.length > 1 ? digits[digits.length - 1] : firstDigit;

  return Number(firstDigit! + lastDigit!);
};

fs.readFile("./input.txt", "utf8", (err, data) => {
  if (err) {
    console.error(`Error in reading file: ${err}`);
    return;
  }

  // console.log(`File Content: ${data}`);

  const lines = data.split("\n");

  const allNumbers = lines.map(getNumberFromLine);

  const sum = allNumbers.reduce((a, b) => a + b, 0);

  console.log(sum);
});
