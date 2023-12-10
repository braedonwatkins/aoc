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

function processData(data: string): any {
  // console.log(`File Content: ${data}`);

  const lines = data.split("\n");

  return partOne(lines);

  partTwo();
}

function partOne(lines: Array<string>) {
  lines.map((line: string) => {
    let gameId: number = Number(line.match(/Game (\d+)/)?.[1] ?? "0");
    console.log(gameId);
    if (gameId === 0) return -1;
  });
}

function partTwo() {}
