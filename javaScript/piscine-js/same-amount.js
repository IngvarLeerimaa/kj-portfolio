/* Instructions
Create a function named sameAmount, that takes three arguments: a string, and 2 regular expressions.Your function should return a boolean.

The objective is to confirm that the regular expressions match the string the same number of times.

 */

function sameAmount(str, re1, re2) {

    let rex = new RegExp(re1, 'g');

    let rex2 = new RegExp(re2, 'g');

    let match1 = str.match(rex);

    let match2 = str.match(rex2);


    if (match1 && match2 && match1.length === match2.length) {
        return true;
    }
    return false;
}