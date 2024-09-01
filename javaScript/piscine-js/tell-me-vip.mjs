/* Instructions
Create a tell-me-vip.mjs script that filters the guests who actually answered 'YES' to your invitation,
 and save this list in a vip.txt file.

The output must print one guest per line, in ascending alphabetic order, 
and formated as following: Number. Lastname Firstname (starting from 1).
 */

import { readdirSync, readFileSync, writeFileSync } from 'fs';


const capitalize = (str) => { return `${str[0].toUpperCase()}${str.slice(1).toLowerCase()}`}

const path = process.argv[2]



try {
  const files = readdirSync(path);
  let names = files.map ( (file) => { 
    const obj = JSON.parse(readFileSync(`${path}/${file}`, 'utf8'))

    if (obj.answer === 'yes') {
      const name = file.slice(0, -5).split('_')
      return `${capitalize(name[1])} ${capitalize(name[0])}`
    }
      
  }).sort()

  names = names.filter( (el) => el !== undefined )

  names = names.map ( (el, i) => `${i+1}. ${el}` )
  
  writeFileSync('vip.txt', names.join("\n"))
      
} catch (err) {
  console.error(err);
}