/* Subject
Create a function named format which accepts a valid Date and a format string. Your function should return a correctly formatted string.

Your function must handle:

y
yyyy
G
GGGG
M
MM
MMM
MMMM
d
dd
E
EEEE
h
hh
m
mm
s
ss
H
HH
a
 */

function format(date, format) {
    console.log(date)

    switch (format) {
        case 'y':
            return Math.abs(date.getFullYear()).toString();

        case 'yyyy':
            let res = "";
            if (date.getFullYear() < 10 && date.getFullYear() > 0) {
                return '000' + Math.abs(date.getFullYear()).toString();
            } else if( date.getFullYear() < 100 && date.getFullYear() > 0) {
                return Math.abs(date.getFullYear()).toString();
            } else if( date.getFullYear() < 1000 && date.getFullYear() > 0) {
                return '0' + Math.abs(date.getFullYear()).toString();
            } else {
                return Math.abs(date.getFullYear()).toString();
            }
            


        case 'G':
             
             if (date.getFullYear() < 0) {
                return 'BC';
            } else {
                return 'AD';
            }



        case 'GGGG':
           if (date.getFullYear() < 0) {
                return 'Anno Domini';
            } else {
                return 'Before Christ';
            }

        case 'M':
            return date.getMonth() + 1;

        case 'MM': 
        if (date.getMonth() + 1 < 10) {
                return '0' + (date.getMonth() + 1);
            }
            return date.getMonth() + 1;
        
        case 'MMM':

        return date.toLocaleString('default', { month: 'short' });
        /* 
            let month = date.getMonth() + 1;
            switch (month) {
                    case 1:
                    return 'Jan';
            
                    case 2:
                    return 'Feb';

                    case 3:
                    return 'Mar';

                    case 4:
                    return 'Apr';

                    case 5:
                    return 'May';

                    case 6:
                    return 'Jun';

                    case 7:
                    return 'Jul';

                    case 8:
                    return 'Aug';

                    case 9:
                    return 'Sep';

                    case 10:
                    return 'Oct';

                    case 11:
                    return 'Nov';

                    case 12:
                    return 'Dec';
            } */

        case 'MMMM':
            return date.toLocaleString('default', { month: 'long' });

        case 'd':
            return date.getDate();

        case 'dd':
            if (date.getDate() < 10) {
                return '0' + date.getDate();
            }
            return date.getDate();

        case 'E':
            return date.toLocaleString('default', { weekday: 'short' });

        case 'EEEE':
            return date.toLocaleString('default', { weekday: 'long' });

        case 'h':
            return date.getHours() % 12;

        case 'hh':
            if (date.getHours() % 12 < 10) {
                return '0' + date.getHours() % 12;
            }
            return date.getHours() % 12;

        case 'm':
            return date.getMinutes();

        case 'mm':
            if (date.getMinutes() < 10) {
                return '0' + date.getMinutes();
            }
            return date.getMinutes();

        case 's':
            return date.getSeconds();

        case 'ss':
            if (date.getSeconds() < 10) {
                return '0' + date.getSeconds();
            }
            return date.getSeconds();

        case 'H':
            return date.getHours();

        case 'HH':
            if (date.getHours() < 10) {
                return '0' + date.getHours();
            }
            return date.getHours();

        case 'a':
            if (date.getHours() < 12) {
                return 'AM';
            }
            return 'PM';

        default:
            return 'Invalid format';
        }





            







  
    
}