/*Create 3 functions:

first: that takes an array or a string and returns its first element or character.

last: that takes an array or a string and return its last element or character.

kiss: that takes an array or string, and returns an array of 2 elements. 
The returned array should contain the last and first elements or characters, in that order.

*/

function first(input) {
    return input[0];
}
    function last(input) {
        return input[input.length - 1];
    }

    function kiss(input) {
        return [last(input), first(input)];
        }