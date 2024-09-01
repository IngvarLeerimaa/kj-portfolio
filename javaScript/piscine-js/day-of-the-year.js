/* Create a function named dayOfTheYear which accepts a Date. 
Your function should return the number of days since the first day of that year. */

function dayOfTheYear(date) {
    let year = date.getFullYear()
    let month = date.getMonth()
    let day = date.getDate()
    let count = 0;
    for (let i = 1; i < month + 1; i++){
        if (i === 2){
            if ((year % 4 === 0 && year % 100 !== 0) || year % 400 === 0){
                count += 29
            } else {
                count += 28
            }
        } else if (i === 4 || i === 6 || i === 9 || i === 11){
            count += 30
        } else {
            count += 31
        }
    }
    return count + day
    

}