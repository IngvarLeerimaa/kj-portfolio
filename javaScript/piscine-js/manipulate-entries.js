/* Instructions
Create 3 functions which work like the .filter, .map and .reduce array methods, but for the entries in the grocery cart.

filterEntries: filters using both key and value.
mapEntries: changes the key, the value or both.
reduceEntries: reduces the entries.
Create 3 additional functions that use your previously created functions and take an object as input:

totalCalories: that will return the total calories of a cart.
lowCarbs: that leaves only those items which are lower than 50 grams.
cartTotal: that will give you the right amount of calories, proteins... and all the other items in your grocery cart.
Think about the shape of Object.entries() */

const filterEntries = (cart, callback) => {
    return Object.fromEntries(Object.entries(cart).filter(callback))
}

const mapEntries = (cart, callback) => {
    return Object.fromEntries(Object.entries(cart).map(callback))
}

function reduceEntries(obj, func, initialValue) {
    if (initialValue == undefined) {
        return Object.entries(obj).reduce(func)
    }
    //not sure why it works
    return Object.entries(obj).reduce(func, initialValue)
}

const totalCalories = (cart) => {
    return Number(reduceEntries(cart, ((acc, [key, value]) =>
    (acc + (nutritionDB[key].calories * value) / 100
    )), 0).toFixed(1))
}

const lowCarbs = (cart) => {


    return filterEntries(cart, ([key, value]) => 
    (nutritionDB[key].carbs * value / 100) < 50)
}

const cartTotal = (cart) => {
    return mapEntries(cart, ([key, value]) => {
        let result = {};
        for (let [k, val] of Object.entries(nutritionDB[key]))
            result[k] = parseFloat(((val * value) / 100).toFixed(3))
        return [key, result]
    })
}