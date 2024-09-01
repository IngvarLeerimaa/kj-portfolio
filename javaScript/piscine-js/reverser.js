function reverse(arr) {

    if (Array.isArray(arr)){
    let newArray = [];
    for (let i = arr.length; i > 0; i--) {
        newArray.push(arr[i - 1]);
    }
    return newArray;
    } else if (typeof arr === "string") {
        let newString = "";
        for (let i = arr.length; i > 0; i--) {
            newString += arr[i - 1];
        }
        return newString;
    }

    return arr;

}
