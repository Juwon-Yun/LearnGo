/* get input */
const fs = require("fs");
const input = (
  process.platform === "linux"
    ? fs.readFileSync("/dev/stdin").toString()
    : 
`input`
).split("\n");
/* get input end */

/* set param */
const n = Number(input.shift());

/* set param end */

/* solve */
const solution = () => {

};
/* solve end */

/* print output */
const output = solution();
console.log(output);
/* print output end */