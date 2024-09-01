import { readdirSync } from 'fs';

const path = process.argv[2]
//350
try {
  const files = readdirSync(path);
  console.log(files.length)
} catch (err) {
  console.error(err);
}