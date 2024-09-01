/* Instructions
Create a map function that takes an array as the first argument, a function as second,
 and that works like the method .map

Create a flatMap function that takes an array as the first argument, a function as second,
 and that works like the method .flatMap 
 
 Code provided
The provided code will be added to your solution, and does not need to be submitted.

Array.prototype.map = undefined
Array.prototype.flatMap = undefined
Array.prototype.flat = undefined*/

const map = (arr, func) => {

    let newArray = [];
    for (let i = 0; i < arr.length; i++){
       newArray.push(func(arr[i], i, arr))
    }
    return newArray;
}

function flatMap(arr, action) {
    return arr.reduce(
        (acc, val, i, arr) => acc.concat(action(val, i, arr)),
        []
    );
}