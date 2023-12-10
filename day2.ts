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

/*************** BOILERPLATE END ***************/

type ColorCount = {
  "red": number;
  "green": number;
  "blue": number;
};
const colorCount: ColorCount = {
  "red": 0,
  "green": 0,
  "blue": 0,
};

function partOne(lines: Array<string>) {
  lines.map((line: string) => {
    let idGameSplit = line.split(":");
    let gameId: number = Number(idGameSplit[0].match(/\d+/));
    // console.log(gameId);

    let game = idGameSplit[1];
    // console.log(counts);

    let rounds = game.split(";");
    // console.log(rounds);

    rounds.map((round) => {
      let counts = round.split(",");
      console.log(counts);
    });
  });
}

function partTwo() {}
