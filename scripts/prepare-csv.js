const fs = require("fs");
const readline = require("readline/promises");
const cidrCalc = require("cidr-calc");
const { gzipSync } = require("zlib");

const reader = fs.createReadStream("./dbip.csv");
const data = [];

readline.createInterface(reader).on("line", (line) => {
  const fields = line.split(',');
  const ip1 = new cidrCalc.Ipv4Address.of(fields[0]);
  const ip2 = new cidrCalc.Ipv4Address.of(fields[1]);
  const ipRange = new cidrCalc.IpRange(ip1, ip2);
  const usageType = fields[17];
  if(usageType !== 'hosting') return;
  ipRange.toCidrs().map(item => item.toString()).forEach(item => data.push(item));
}).on('close', () => {
  const result = gzipSync(JSON.stringify(data));
  fs.writeFileSync('ips.json.gz', result);
});
