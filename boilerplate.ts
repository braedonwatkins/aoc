import * as fs from "fs";

const filePath = "./input.txt";

fs.readFile("./input.txt", "utf8", (err, data) => {
  if (err) {
    console.error(`Error in reading file: ${err}`);
    return;
  }
  console.log(`File Content: ${data}`);
});
