/* structions
Create the following functions:

Your solutions must use filter.

filterShortStateName: accepts an array of strings, and returns only those strings which contain less than 7 characters.

filterStartVowel: accepts an array of strings, and returns only those that start with any vowel (a,e,i,o,u).

filter5Vowels: accepts an array of strings, and returns only those which contain at least 5 of any vowels (a,e,i,o,u).

filter1DistinctVowel: accepts an array of strings, and returns only those which contain distinct vowels (a,e,i,o,u). For example, "Alabama" contains only 1 distinct vowel "a".

multiFilter: accepts an array of objects, and returns only those which:

the key capital contains at least 8 characters.
the key name does not start with a vowel.
the key tag has at least one vowel.
the key region is not "South"
 */

const filterShortStateName = (arr) => arr.filter((onPieceOfIt) =>
    onPieceOfIt.length < 7);
    
const filterStartVowel = (arr) => arr.filter((element) =>
    element[0] === "A" || element[0] === "E" || element[0] === "I" || element[0] === "O" || element[0] === "U");

const filter5Vowels = (arr) => arr.filter((randomName) =>
    randomName.match(/[aeiou]/gi).length > 4
); 

const filter1DistinctVowel = (arr) => arr.filter((item) => new Set(item.toLowerCase().match(/[aeiou]/gi)).size === 1
);

const multiFilter = (arr) => arr.filter((item) => {
    let capital = item.capital.length >= 8;
    let name = !/^[aeiou]/i.test(item.name);
    let tag = /[aeiou]/i.test(item.tag);
    let region = item.region !== "South";
    return capital && name && tag && region;
});

