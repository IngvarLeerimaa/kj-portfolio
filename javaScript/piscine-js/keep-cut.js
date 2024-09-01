function strToArr(str) {
    return str.split('');
}

function cutFirst(str){
    str = strToArr(str)
    return str.slice(2).join('')

}
function cutLast(str){
    let x = str.length;
    str = strToArr(str);
    return str.slice(0, x-2).join('')
}
function cutFirstLast(str){
    let x = str.length;
    str = strToArr(str);
    return str.slice(2, x-2).join('')
}
function keepFirst(str){
    str = strToArr(str);
    return str.slice(0, 2).join('')
}
function keepLast(str){
    let x = str.length;
    str = strToArr(str);
    return str.slice(x-2, x).join('')
}
function keepFirstLast(str){
    if (str.length < 4) {
        return str;
    }
    let x = str.length;
    str = strToArr(str);
    return str.slice(0, 2).join('') + str.slice(x-2, x).join('')
}
