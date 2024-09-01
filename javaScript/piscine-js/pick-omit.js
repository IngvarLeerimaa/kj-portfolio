/* Create two functions which takes an object and a string or array of strings. 

They should return a new object which:

pick: contains only those keys which appear in the string or array of strings.
omit: contains only those keys which do not match the string, or do not appear in the array of strings.

Those functions are pure and must not modify the given object */

const pick = (obj, keys) => {
    const res = {};
    if (!Array.isArray(keys)) keys = Array(keys);
    for (let element of keys) {
      if (obj.hasOwnProperty(element)) {
        res[element] = obj[element];
      }
    }
    return res;
  };
  const omit = (obj, keys) => {
    const res = {
      ...obj,
    };
    if (!Array.isArray(keys)) keys = Array(keys);
    for (let element of keys) {
      if (obj.hasOwnProperty(element)) {
        delete res[element];
      }
    }
    return res;
  };