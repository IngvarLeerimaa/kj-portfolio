/* Instructions
Create a function named forEach which takes an array as the first argument, 
a function as the second argument, and that works like the 
Array.prototype.forEach method.
 */
/* 
const forEach = (arr, func) => arr.map(func);
 */

/*Problem : Test is passable with using maps.  
I might be wrong but Array.prototype.forEach() method executes a provided 
function once for each array element, and does not return a new array.  
Map creates a new array rather than performing a side effect on each element of the array.  
*/

const forEach = (arr, func) => {
    for (let i = 0; i < arr.length; i++) {
      func(arr[i], i, arr);
    }
  };