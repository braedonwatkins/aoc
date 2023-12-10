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

const getNumberFromLine = (line: string) => {
  const spelledNumsStr = `/one|two|three|four|five|six|seven|eight|nine/`;
  const numsStr = `/\d/`;

  const regex = new RegExp(
    `one|two|three|four|five|six|seven|eight|nine|\\d`,
    "g"
  );

  let match;
  let firstMatch: string = "-1";
  let lastMatch: string = "-1";

  while ((match = regex.exec(line)) !== null) {
    if (firstMatch === "-1") {
      firstMatch = match[0];
    }
    console.log(regex);
    lastMatch = match[0];

    regex.lastIndex = match.index + 1;
  }

  console.log(firstMatch, lastMatch);
  if (firstMatch in numStrObj) firstMatch = numStrObj[firstMatch];
  if (lastMatch in numStrObj) lastMatch = numStrObj[lastMatch];

  return Number(firstMatch! + lastMatch!);
};

function processData(data: string): any {
  // console.log(`File Content: ${data}`);

  const lines = data.split("\n");

  const allNumbers = lines.map(getNumberFromLine);
  // console.log(allNumbers);

  const sum = allNumbers.reduce((a, b) => a + b, 0);

  return sum;
}
