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

// DAY 2

type NumStrObjType = {
  [key: string]: string;
  one: string;
  two: string;
  three: string;
  four: string;
  five: string;
  six: string;
  seven: string;
  eight: string;
  nine: string;
};

const numStrObj: NumStrObjType = {
  "one": "1",
  "two": "2",
  "three": "3",
  "four": "4",
  "five": "5",
  "six": "6",
  "seven": "7",
  "eight": "8",
  "nine": "9",
};

const getNumberFromLine2 = (line: string) => {
  const regex = new RegExp(
    `one|two|three|four|five|six|seven|eight|nine|\\d`,
    "g"
  );

  let match;
  let firstMatch: string = "0";
  let lastMatch: string = "0";

  while ((match = regex.exec(line)) !== null) {
    if (firstMatch === "0") {
      firstMatch = match[0];
    }
    lastMatch = match[0];

    regex.lastIndex = match.index + 1;
  }

  if (firstMatch in numStrObj) firstMatch = numStrObj[firstMatch];
  if (lastMatch in numStrObj) lastMatch = numStrObj[lastMatch];

  console.log(firstMatch, lastMatch);
  return Number(firstMatch! + lastMatch!);
};

function processData(data: string): number {
  // console.log(`File Content: ${data}`);

  const lines = data.split("\n");

  const allNumbers = lines.map(getNumberFromLine);
  // console.log(allNumbers);

  const sum = allNumbers.reduce((a, b) => a + b, 0);

  return sum;
}
