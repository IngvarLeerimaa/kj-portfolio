function nasa(N) {
    var s = "";
    for (var i = 1; i <= N; i++) {
        switch (true) {
            case (i % 3 == 0 && i % 5 == 0):
                
                s += "NASA";

                if (i != N) {
                    s += " ";
                }

                break;
            case (i % 3 == 0):
                s += "NA";
                if (i != N) {
                    s += " ";
                }
                break;
            case (i % 5 == 0):
                s += "SA";
                if (i != N) {
                    s += " ";
                }
                break;
            default:
                s += i.toString();
                if (i != N) {
                    s += " ";
                }
                break;
        }
    }
    return s;
}
