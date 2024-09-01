/* Instructions
I do not want onions. I want oranges.

Create 3 functions that works like the .filter, .map and .reduce array methods,
but for the keys of your grocery cart. You can see their names and how they work in the examples.
 */

const filterKeys = (cart, callback) => {
    return Object.fromEntries(Object.entries(cart).filter(([key, _]) => callback(key)))
}

const mapKeys = (cart, callback) => {
    return Object.fromEntries(Object.entries(cart).map(([key, value]) => [callback(key), value]))
}


function reduceKeys(obj, func, initialValue) {
    if (initialValue == undefined) {
        return Object.keys(obj).reduce(func)
    }
    //not sure why it works
    return [initialValue].concat(Object.keys(obj)).reduce(func)
}