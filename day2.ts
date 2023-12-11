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
type Color = keyof ColorCount;

const colorCount: ColorCount = {
  "red": 0,
  "green": 0,
  "blue": 0,
};

const hardCoded: ColorCount = {
  "red": 12,
  "green": 13,
  "blue": 14,
};

const winningRounds: Array<number> = [];

function partOne(lines: Array<string>) {
  let totalScore = 0;

  lines.forEach((line: string) => {
    let [gameIdStr, game] = line.split(": ");
    let gameId = parseInt(gameIdStr.replace("Game ", ""));
    let isWinningGame = true;

    let rounds = game.split("; ");
    rounds.forEach((round) => {
      if (!isWinningGame) return;

      let counts = round.split(", ");
      let currentColorCount = { ...colorCount };

      for (let count of counts) {
        let [numStr, color] = count.split(" ");
        let num = parseInt(numStr);
        let colorKey = color as Color; // Asserting the type

        currentColorCount[colorKey] += num;
        if (currentColorCount[colorKey] > hardCoded[colorKey]) {
          isWinningGame = false;
          break;
        }
      }
    });

    if (isWinningGame) {
      totalScore += gameId;
    }
  });

  return totalScore;
}

function partTwo() {}
