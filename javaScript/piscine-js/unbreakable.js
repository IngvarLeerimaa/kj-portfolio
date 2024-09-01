function split(str, separator) {
    // Split a given string using a multi-character separatorarator
    // and return an array of the results.
    if (separator === null) {
        separator = ",";
    }
    var result = [];
    if (separator === "") {
        for (var i = 0; i < str.length; i++) {
            result.push(str[i]);
        }
        return result;
    }
    var end = str.indexOf(separator);
    while (end > -1) {
        end = str.indexOf(separator);
        if (end === -1) {
            break;
        }
        result.push(str.slice(0, end));
        str = str.slice(end + separator.length);
    }
    result.push(str);
    return result;
}

function join(arr, separator) {
    if (separator === null) {
        separator = ",";
    }
    var result = arr[0].toString();
    for (var i = 1; i < arr.length; i++) {
        result += separator + arr[i];
    }
    return result;
}
