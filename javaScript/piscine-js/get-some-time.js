/* Instructions
Create a function named firstDayWeek, which accepts a specific week in a given year:

number: representing a week of the year (between 1 and 53).
string: representing a year.
Your function should return a string representing the date of the first day of that specific week in the format dd-mm-yyyy.

Week 1 is in all cases, the week which contains the 1st of January.

The first day of a week is a date.

If the start of the week is in the previous year, then your function should return the first day of the specified year.
 */

function firstDayWeek(week, year) {

    console.log(week)
    console.log(year)


    let days = 1 + ((week-1) *7)
    console.log(days)
   
    let date = new Date(year, 0, days);
    console.log(date)


    while (date.getDay() !== 1) {
        if (date.getFullYear() == year - 1) {
            return '01-01-' + year.toString()
        }
        date.setDate(date.getDate() - 1);
    }
    
    if (year.toString().slice(0, 2) == '00') {
        date.setDate(date.getDate() + 1)
    }







    let month = date.getMonth()

    let thisYear = date.getFullYear()

    let day = date.getDate()

    if (year.toString().slice(0, 2) == '00') {
        thisYear = "00" + date.getFullYear().toString().slice(-2)
    }


//printing
    let res = "";

    if (day < 10) {
        res = "0" + day.toString();
    } else {
        res = day;
    }

    if (month +1 < 10) {
        res = res + "-0" + (month+1).toString();
    } else {
        res = res + "-" + (month+1).toString();
    }

    res = res + "-" + thisYear;
    return res;

}