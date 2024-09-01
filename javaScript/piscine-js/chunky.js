function chunk (arr, size) {
  var chunked = [];
  for (var i = 0; i < arr.length; i += size) {
    chunked.push(arr.slice(i, i + size));
  }
  return chunked;
}