/* Instructions
Create a function named ionOut, that receives a string and returns an array with every word
 containing 'ion' following a 't'.The words should be returned without the 'ion' part.
 */


function ionOut(str) {
    let arr = str.split(" ");
    let rex = /tion/g;
    let rex2 = /[.,?!]/g;
    let res = [];
    arr.forEach((word) => {
        word.match(rex) ? res.push(word.replace(rex2, "").slice(0, -3)) : null;
    });
    return res;
}