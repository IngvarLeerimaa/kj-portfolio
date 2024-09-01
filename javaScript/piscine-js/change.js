function get(str){
    return sourceObject[str.toString()]

};

function set(str, value){
    sourceObject[str.toString()] = value
    return sourceObject[str.toString()]

};
