/* Instructions
Create the function named flow that will act like the _.flow([funcs]) from lodash.

Example
const square = (nbr) => nbr * nbr
const add2Numbers = (nbr1, nbr2) => nbr1 + nbr2

const flowedFunctions = flow([add2Numbers, square])
flowedFunctions(2, 3) // -> 25
 */

const flow = (funcs) => {
    console.log(funcs);
    return (args) => {
        return funcs.reduce((acc, func) => {
            return func(acc);
        }, args);
    }
}