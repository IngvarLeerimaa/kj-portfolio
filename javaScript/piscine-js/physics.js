function getAcceleration(x) {
    if (x.f !== undefined && x.m !== undefined) {
        return x.f / x.m;
    } else if (x.Δv !== undefined && x.Δt !== undefined) {
        return x.Δv / x.Δt;
    } else if (x.d !== undefined && x.t !== undefined) {
        return (2 * x.d) / (x.t ** 2);
    } else {
        return 'impossible';
    }
}


/*
a = F/m
a = Δv/Δt
a = 2d/t^2
possible to use hasOwnProperty()

*/