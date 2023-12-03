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

function processData(data: string): number {
  const lines = data.split("\n");
  const allNumbers = lines.map(getNumberFromLine);
  const sum = allNumbers.reduce((a, b) => a + b, 0);
  return sum;
}

const getNumberFromLine = (line: string) => {
  const digits = line.match(/\d/g) || [];
  if (digits.length === 0) return 0;

  const firstDigit = digits[0];
  const lastDigit = digits.length > 1 ? digits[digits.length - 1] : firstDigit;

  return Number(firstDigit! + lastDigit!);
};

async function main() {
  const filePath = process.argv[2] || "input.txt"; // node file.ts input.txt

  try {
    const fileContent = await readFileContent(filePath);
    const result = processData(fileContent);
    console.log(result);
  } catch (err) {
    console.error("Error occurred:", err);
  }
}

main();
