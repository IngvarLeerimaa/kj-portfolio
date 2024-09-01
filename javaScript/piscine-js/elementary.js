function multiply(a,b){
   let buf = 0;

   if (a === 0 || b === 0){
         return 0;
    } else {
        if (b < 0){
            a = -a;
            b = -b;
        }
        buf = a + multiply(a, b-1);
    }

    return buf;

}
function divide(a,b){
    let buf = 0;
    let isNegative = false;

    if (a == 0){
        buf = 0;
    } else if (b==0){
        //start
        if (a>0){
            buf = Infinity;
        } else {
            buf = -Infinity;
        }
        //end
    } else if (b ==1){
        buf = a;
    } else if (b == -1){
        buf = -a;
    } else {
        if (a < 0){
            a = Math.abs(a)
            isNegative = true;
        }

        if(b<0){
            b = Math.abs(b)
            if (isNegative){
                isNegative = false;
            } else {
                isNegative = true;
            }
        }

   
    while (a>=b){
        a-=b;
        buf++;
    }

    if (isNegative){
        buf = -buf;
    }

}
return buf;
}

function modulo(a,b){
    return a - multiply(b, divide(a,b));
}