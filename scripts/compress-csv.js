const fs = require("fs");
const { gzipSync } = require("zlib");
const contents = fs.readFileSync('./prepared.csv');
const compressed = gzipSync(contents);
fs.writeFileSync('prepared.csv.gz', compressed);