/* nstructions
Create a function named letterSpaceNumber that accepts a string; 
returning an array with every instance of a letter, 
followed by a space, followed by a number, 
only if that number has only one digit, and is not followed by any letter.

    Examples
console.log(letterSpaceNumber('example 1, example 20'))
// output: ['e 1']
 */

function letterSpaceNumber(str){
    
    
    let regex = /[A-Za-z]\s\d(?![a-z0-9])/gi;
    let result = str.match(regex);


    return result !== null ? result : [];


}