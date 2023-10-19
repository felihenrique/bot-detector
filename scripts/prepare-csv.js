const fs = require("fs");
const readline = require("readline/promises");
const cidrCalc = require("cidr-calc");

const writer = fs.createWriteStream("./prepared.csv");

const reader = fs.createReadStream("./dbip.csv");
readline.createInterface(reader).on("line", (line) => {
  const fields = line.split(',');
  const ip1 = new cidrCalc.Ipv4Address.of(fields[0]);
  const ip2 = new cidrCalc.Ipv4Address.of(fields[1]);
  const ipRange = new cidrCalc.IpRange(ip1, ip2);
  const usageType = fields[17];
  const types = ['hosting'];
  if(!types.includes(usageType)) return;
  writer.write(JSON.stringify(ipRange.toCidrs().map(item => item.toString())) + "\n");
});
