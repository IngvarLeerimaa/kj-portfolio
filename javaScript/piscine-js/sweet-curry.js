/* Instructions
Create the following functions with the "currying" process.
 Those functions should accept only one argument each.

mult2: that multiplies two numbers.
add3: that adds three numbers.
sub4: that subtracts four numbers.
Notions
 */


const mult2 = (a) => (b) => a * b;

const add3 = (a) => (b) => (c) => a + b + c;

const sub4 = (a) => (b) => (c) => (d) => a - b - c - d;