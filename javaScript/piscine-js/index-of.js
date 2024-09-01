

function indexOf (array, value, fromIndex) {
    if(fromIndex === undefined) {
        fromIndex = 0
    }

    let result;

    for(let i = fromIndex; i < array.length; i++) {
        if(array[i] === value) {
            result = i;
            return result;
        }     
    }

    if (result === undefined) {
        return -1;
    } 
}

function lastIndexOf (array, value, fromIndex) {
    if (fromIndex === undefined) {
        fromIndex = array.length - 1
    }

    let result;

    for (let i = fromIndex; i >= 0; i--) {
        if (array[i] === value) {
            result = i;
            break;
        }
    }

    if (result === undefined) {
        return -1;
    }
    return result;
}

function includes (array, value, fromIndex) {
    if (fromIndex === undefined) {
        fromIndex = 0
    }

    let result;

    for (let i = fromIndex; i < array.length; i++) {
        if (array[i] === value) {
            result = true;
            return result;
        }
    }

    if (result === undefined) {
        return false;
    }
}
