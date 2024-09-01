/* Instructions
Create a function named all that works like Promise.all but with objects (instead of arrays).

Code provided
The provided code will be added to your solution, and does not need to be submitted.

Promise.all = undefined
*/

function all(object) {
    return new Promise(async (resolve, reject) => {
        let result = {}
        for (const key in object) {
            try {
                result[key] = await object[key]
            } catch (e) {
                reject(e)
            }
        }
        resolve(result)
    })
}