/* Instructions
Create these functions which receive an array and a function each. Each element will return true if

every: every element in the array respects the condition of the function.
some: that returns true if at least one element in the array respects the condition of the function.
none: that returns true if none of the elements in the array respects the condition of the function.
The use of [].every and [].some is forbidden for this exercise.

Code provided
The provided code will be added to your solution, and does not need to be submitted.

Array.prototype.some = Array.prototype.every = undefined
 */

const every = (array, fn) => {
      for (let i = 0; i < array.length; i++) {
    if (!fn(array[i])) {
      return false;
    }
  }
  return true;
}

const some = (array, fn) => {
    for (let i = 0; i < array.length; i++) {
        if (fn(array[i])) {
        return true;
        }
    }
    return false;
    }

const none = (array, fn) => {
    for (let i = 0; i < array.length; i++) {
        if (fn(array[i])) {
        return false;
        }
    }
    return true;
    }

