/*
arrToSet: from Array to Set.
arrToStr: from Array to string.
setToArr: from Set to Array.
setToStr: from Set to string.
strToArr: from string to Array.
strToSet: from string to Set.
mapToObj: from Map to Object.
objToArr: from Object to Array.
objToMap: from Object to Map.
arrToObj: from Array to Object.
strToObj: from string to Object.
*/
function arrToSet(arr) {
    return new Set(arr);
}
function arrToStr(arr) {

    return arr.join('');
}
function setToArr(str) {
    return Array.from(str);
}
function setToStr(set) {
    return Array.from(set).join('');
}
function strToArr(str) {
    return str.split('');
}
function strToSet(str) {
    return new Set(str.split(''));
}
function mapToObj(map) {
    return Object.fromEntries(map);
}
function objToArr(obj) {
    return Object.values(obj);
}
function objToMap(obj) {
    return new Map(Object.entries(obj));
}
function arrToObj(arr) {
    return Object.assign({}, arr);
}
function strToObj(str) {
    return Object.assign({}, str.split(""));
}


function superTypeOf(x) {
    if (typeof x === 'string') {
        return "String";
    } else if (typeof x === 'number') {
        return "Number";
    } else if (typeof x === 'boolean') {
        return "boolean";
    } else if (x === undefined) {
        return "undefined";
    } else if (x === null) {
        return "null";
    }  else if (Array.isArray(x)) {
        return "Array";
    } else if (x instanceof Set) {
        return "Set";
    } else if (x instanceof Map) {
        return "Map";
    } else if (typeof x === 'object') {
        return "Object";
    } else if (typeof x === 'function') {
        return "Function";
    } else {
        return "unknown";
    }
}
