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
    console.log(result);
  } catch (err) {
    console.error("Error occurred:", err);
  }
}

main();

/************** END BOILERPLATE  *****************/

function processData(data: string): any {
  const lines = data.split("\n");
}
