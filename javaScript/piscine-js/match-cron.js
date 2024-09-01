/* Instructions
Create a function named matchCron, which accepts a valid cron string, and a valid Date.
 Your function should return true if the pattern matches the Date.

You only need to implement numbers and *. Other complex patterns are not required.

Only valid patterns will be tested.

Example
matchCron('9 * * * *', new Date('2020-05-30 18:09:00')) // -> true
matchCron('9 * * * *', new Date('2020-05-30 19:09:00')) // -> true
matchCron('9 * * * *', new Date('2020-05-30 19:21:00')) // -> false
//         | | | | |
//         | | | | +- Day of the Week   (range: 1-7, 1 is Monday)
//         | | | +--- Month of the Year (range: 1-12)
//         | | +----- Day of the Month  (range: 1-31)
//         | +------- Hour              (range: 0-23)
//         +--------- Minute            (range: 0-59) */

function matchCron(cron, date) {
  const [minute, hour, day, month, dayOfWeek] = cron.split(' ')
  const dateMinute = date.getMinutes()
  const dateHour = date.getHours()
  const dateDay = date.getDate()
  const dateMonth = date.getMonth() + 1
  const dateDayOfWeek = date.getDay() === 0 ? 7 : date.getDay()

  return (
    (minute === '*' || minute === dateMinute.toString()) &&
    (hour === '*' || hour === dateHour.toString()) &&
    (day === '*' || day === dateDay.toString()) &&
    (month === '*' || month === dateMonth.toString()) &&
    (dayOfWeek === '*' || dayOfWeek === dateDayOfWeek.toString())
  )
}