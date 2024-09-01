const round = (n) => {
    switch (n) {
        case Number.NaN:
            return 0;
        case n <= Number.MIN_VALUE:
            return Number.MIN_VALUE;
        case n >= Number.MAX_VALUE:
            return Number.MAX_VALUE;
        default:
            let r = modulo(n, 1);
            if (r == 0) {
                return n;
            } else {
                if (n > 0) {
                    if (r < 0.5) {
                        return n - r;
                    } else {
                        return n + (1 - r);
                    }
                } else {
                    if (r > -0.5) {
                        return n - r;
                    } else {
                        return n - (1 + r);
                    }
                }
            }
    }
};

const ceil = (n) => {
    return -floor(-n);
};

const floor = (n) => {
    if (Number.NaN) {
        return 0;
    }
    let r = modulo(n, 1);
    if (r == 0) {
        return n;
    } else {
        if (r < 0) {
            return n - (1 + r)
        } else {
            return n - r;
        }

    }
};

const trunc = (n) => {
    if (n >= 0xfffffffff) {
        let res = n - 0xfffffffff
        return trunc(res) + 0xfffffffff
    }

    if (n >= 0) {
        return floor(n);
    } else {
        return ceil(n);
    }
};

const modulo = (a, b) => {
    const aSign = Math.sign(a);

    a = Math.abs(a);
    b = Math.abs(b);
    var res = 0;

    let i = 0;
    while (i <= a) {
        i += b;
    }


    i -= b;

    res = a - i;
    if (aSign < 0) {
        res = -res;
    }

    return parseFloat(res);
};