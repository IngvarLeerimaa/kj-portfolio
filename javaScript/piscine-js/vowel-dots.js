/* Instructions
Create a function named vowelDots that receives a string.Your function should return a new string with a.after every vowel.

Your RegEx should be stored in a variable named vowels.

    a, e, i, o and u are considered as vowels here.

 */

const vowels = /[aeiou]/gi;
function vowelDots(str) {
  
  
    var result = str.replace(vowels, '$&.');
/* 
    console.log(str)
    console.log(result) */
    return result;

}