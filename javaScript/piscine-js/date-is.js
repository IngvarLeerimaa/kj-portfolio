/* Create the following functions:

isValid: accepts a Date, and returns false if the Date is invalid.
isAfter: accepts two Date arguments, and returns true if the first is greater then the second.
isBefore: accepts two Date arguments, and returns true if the second is greater than the first.
isFuture: accepts a Date, and returns true if the Date is valid, and is after than the present date.
isPast: accepts a Date, and returns true if the Date is valid, and is before the present date.
 */

function isValid(date) {
   if(new Date(date).toString() === 'Invalid Date') {
       return false
   }
   if (!(date instanceof Date) && typeof date !== "number"){
         return false
   }
    return true
    
}

function isAfter(date1, date2) {
    if(isValid(date1) && isValid(date2)) {
        return date1 > date2
    }
    return false
}

function isBefore(date1, date2) {
    if(isValid(date1) && isValid(date2)) {
        return date1 < date2
    }
    return false
}

function isFuture(date) {
    if(isValid(date)) {
        return date > new Date()
    }
    return false
}

function isPast(date) {
    if(isValid(date)) {
        return date < new Date()
    }
    return false
}