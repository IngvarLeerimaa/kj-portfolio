/* Let's eliminate Sundays by taking them out of the calendar,
making a week only six days long, from "Monday" to "Saturday".

Create a function called sunnySunday that takes a Date as an argument and returns the weekday as a string.

01/01/0001 is a Monday. */

function sunnySunday(date) { 
    let day = date.getTime() + 62135596800000;
    const week = {
        0: "Monday",
        1: "Tuesday",
        2: "Wednesday",
        3: "Thursday",
        4: "Friday",
        5: "Saturday",
    };
    return week[(day / 86400000) % 6]; 
}