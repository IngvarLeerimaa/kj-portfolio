function addWeek(date) {

    let day = date.getTime() + 62135596800000;


    const secondWeek = {
        0: "Monday",
        1: "Tuesday",
        2: "Wednesday",
        3: "Thursday",
        4: "Friday",
        5: "Saturday",
        6: "Sunday",
        7: "secondMonday",
        8: "secondTuesday",
        9: "secondWednesday",
        10: "secondThursday",
        11: "secondFriday",
        12: "secondSaturday",
        13: "secondSunday",
    };
    
 
    return secondWeek[(day / 86400000) % 14];
}


function timeTravel({date, hour, minute, second}){
    let newDate = new Date(date);
    newDate.setHours(hour);
    newDate.setMinutes(minute);
    newDate.setSeconds(second);
    return newDate;
}

/* console.log(addWeek(new Date('0001-01-09'))); */