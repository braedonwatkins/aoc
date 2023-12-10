import * as fs from "fs";

async function readFileContent(filePath: string): Promise<string> {
  try {
    const data = await fs.promises.readFile(filePath, "utf8");
    return data;
  } catch (err) {
    console.error(`Error in reading file: ${err}`);
    throw err;
  }
}

async function main() {
  const filePath = process.argv[2] || "input.txt"; // node file.ts input.txt

  try {
    const fileContent = await readFileContent(filePath);
    const result = processData(fileContent);
    console.log("RESULT:", result);
  } catch (err) {
    console.error("Error occurred:", err);
  }
}

main();

/************** END BOILERPLATE  *****************/

type NumStrObjType = {
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

const getNumberFromLine = (line: string) => {
  // console.log(matches);

  const matches = [];
  return 0;
  if (matches.length === 0) return 0;

  let firstDigit: string = matches[0]!;
  let lastDigit: string =
    matches.length > 1 ? matches[matches.length - 1]! : firstDigit!;

  if (firstDigit in numStrObj)
    firstDigit = numStrObj[firstDigit as keyof NumStrObjType];
  if (lastDigit in numStrObj)
    lastDigit = numStrObj[lastDigit as keyof NumStrObjType];

  console.log(firstDigit, lastDigit);
  return Number(firstDigit + lastDigit);
};

function processData(data: string): any {
  // console.log(`File Content: ${data}`);

  const lines = data.split("\n");

  const allNumbers = lines.map(getNumberFromLine);
  // console.log(allNumbers);

  const sum = allNumbers.reduce((a, b) => a + b, 0);

  return sum;
}
