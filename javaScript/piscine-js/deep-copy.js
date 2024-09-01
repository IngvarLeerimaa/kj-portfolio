/* Instructions
Create a function named deepCopy that copies objects and arrays recursively.
 */
function deepCopy(obj) {
    var copy = {};
    if (Array.isArray(obj)) {
        return deepClone(obj)
    }
    if (null == obj || "object" != typeof obj) return obj;
   
    for (var x in obj) {
        if (obj.hasOwnProperty(x)) copy[x] = obj[x];
    }
    return copy;
}

function deepClone(arr) {
    return arr.map(element => Array.isArray(element) ? deepClone(element) : element);
}