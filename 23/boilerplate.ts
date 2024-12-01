import * as fs from "fs";

async function readFileContent(filePath: string): Promise<Array<string>> {
  try {
    const data = await fs.promises.readFile(filePath, "utf8");
    const lines = data.split("\n");
    return lines;
  } catch (err) {
    console.error(`Error in reading file: ${err}`);
    throw err;
  }
}

async function main() {
  const filePath = process.argv[2] || "input.txt"; // node file.ts input.txt

  try {
    const lines = await readFileContent(filePath);
    const result = partOne(lines);
    // const result = partTwo(lines);

    console.log("RESULT:", result);
  } catch (err) {
    console.error("Error occurred:", err);
  }
}

main();

/************** END BOILERPLATE *****************/

const partOne = (lines: Array<string>) => {};

const partTwo = (lines: Array<string>) => {};
