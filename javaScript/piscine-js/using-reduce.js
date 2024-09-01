/* Instructions
Create the following functions:

Your solutions must use reduce.

adder: accepts an array of numbers, and returns the sum as a number.

sumOrMul: accepts an array of numbers and adds or multiplies its elements depending on 
whether the element is odd or even. Even = multiply. Odd = add.

funcExec: accepts an array of functions and executes them using reduce, returning the result.

Each function may accept an optional argument, which should be the initial value for the function's execution.

Example:
sumOrMul([1, 2, 3, 5, 8], 5) // (((((5 + 1) * 2) + 3) + 5) * 8) -> 160
 */

const adder = (arr, maybeMaybeNot) => 
    maybeMaybeNot !== undefined ? 
    arr.reduce((sum, item) => sum + item, maybeMaybeNot) :
    arr.reduce((sum, item) => sum + item, 0);

const sumOrMul = (arr, def) => 

    def !== undefined ? 
    arr.reduce((sum, num) => (num % 2 === 0 ? sum *num : sum + num), def):
    arr.reduce((sum, num) => (num % 2 === 0 ? sum *num : sum + num), 0);

const funcExec = (func, x) => func.reduce((z, y) => y(z), x);


