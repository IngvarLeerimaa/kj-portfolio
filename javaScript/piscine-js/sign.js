function sign(x){
    return x > 0? 1 : x < 0? -1 : 0;
}

function sameSign(x,y){
    return sign(x) === sign(y) ? true : false;
}