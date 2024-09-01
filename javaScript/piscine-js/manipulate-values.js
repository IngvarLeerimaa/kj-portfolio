/* Instructions
Let's buy groceries.

You have a grocery cart with some items you need. 
The item's name is the key, and the value will represent nutrition facts per 100 grams.

Create 3 functions that work like the .filter, .map and .reduce array methods, 
for the values in your grocery cart object. You can see their function names and how they work in the examples. */



const filterValues = (cart, callback) => {

    return Object.fromEntries(Object.entries(cart).filter(([_, value]) => callback(value))
    )};

const mapValues = (cart, callback) => {
    return Object.fromEntries(Object.entries(cart).map(([key, value]) => [key, callback(value)])
    )};

const reduceValues = (cart, callback, initialValue) => {
    if (initialValue === undefined) {
        initialValue = 0
    }
   /*  console.log(cart, callback, initialValue)
    console.log(Object.entries(cart).reduce((acc, [_, value]) => callback(acc, value), initialValue))
    */ 
   return Object.entries(cart).reduce((acc, [_, value]) => callback(acc, value), initialValue)
    };
