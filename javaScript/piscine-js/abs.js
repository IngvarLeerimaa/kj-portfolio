Math.abs = undefined

function isPositive(num) {
    return Number.isInteger(num) && num > 0;
}

//shorter is better
function abs(num) {
  if (isNaN(num)) {
    return NaN;
  }

  if (num === 0) {
    return 0;
}

  return num > 0 ? num : -num;


}

