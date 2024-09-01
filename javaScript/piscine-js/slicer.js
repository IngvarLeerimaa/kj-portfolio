function slice(str, start, end) {
  var result = '';

  if (start < 0) {
    start = str.length + start;
  }

  if (start >= str.length) {
    return result;
  }

  if (end < 0) {
    end = str.length + end;
  }

  if (end < -str.length) {
    end = 0;
  }

  if (end >= str.length || end === undefined) {
    end = str.length;
  }

 if (typeof str === 'string'){
  for (var i = start; i < end; i++) {
    result += str[i];
  
  }
}

  if (Array.isArray(str)){
    result = [];
    for (var i = start; i < end; i++) {
      result.push(str[i]);
    }
  }
  return result;

}
