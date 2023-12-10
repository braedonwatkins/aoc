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

const hardCoded: ColorCount = {
  "red": 12,
  "green": 13,
  "blue": 14,
};

const winningRounds: Array<number> = [];

function partOne(lines: Array<string>) {
  lines.map((line: string) => {
    let idGameSplit = line.split(":");
    let gameId: number = Number(idGameSplit[0].match(/\d+/));
    // console.log(gameId);

    winningRounds.push(gameId); // we will remove if the round loses

    let game = idGameSplit[1];
    // console.log(counts);

    let rounds = game.split(";");
    // console.log(rounds);

    rounds.map((round) => {
      let counts = round.split(",");
      console.log(counts);

      try {
        counts.forEach((count) => {
          let countNum: number = Number(count.match(/\d+/));
          let countColor: string | undefined =
            count.match(/red|blue|green/)?.[0];
          if (!countColor)
            throw new Error(
              `Was unable to extract color from count string, ${count}`
            );

          //TODO: need to do this in a foreach or something where I can bail bc i dont want to pop multiple times
          if (hardCoded[countColor as keyof ColorCount] < countNum) {
            winningRounds.pop(); // removes the "winning" round
            throw new Error("round lost, exit this instance");
          }
        });
      } catch (err) {
        // console.log(err);
        console.log(`round ${gameId} lost, exit this instance`);
      }
    });

    console.log(winningRounds);
  });
}

function partTwo() {}
