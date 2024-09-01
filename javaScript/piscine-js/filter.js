/* Instructions
Create the following functions, which each take an array as the first argument, 
and a function as the second argument.

filter: that works like the [].filter method.

reject: that works like the reject function from lodash.

partition: that works like the partition function from lodash.

Code provided
The provided code will be added to your solution, and does not need to be submitted.

Array.prototype.filter = undefined
 */

const filter = (arr, func) =>{ 
    var result = [];
    for (var i = 0; i < arr.length; i++) {
        if (func(arr[i], i, arr)) {
            result.push(arr[i]);
        }
    }
    return result;
};

const reject = (arr, func) => {
    var result = [];
    for (var i = 0; i < arr.length; i++) {
        if (!func(arr[i], i, arr)) {
            result.push(arr[i]);
        }
    }
    return result;

}

const partition = (arr, func) =>{
    return [filter(arr, func), reject(arr, func)];

}
     