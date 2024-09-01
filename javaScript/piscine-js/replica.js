/* Instructions
Create a function named replica that allows you to deep assign the values of all properties from one or more objects to a obj object.

Watch out for shallow copies.
 */

const isObject = a => typeof a === "object" && !(
    a instanceof Set || 
    a instanceof RegExp || 
    a instanceof Array || 
    a instanceof Map
    );

const replica = (obj, ...other) => {
  other.forEach(ob => {
    Object.entries(ob).forEach(([key, val]) => {
      if (isObject(val) && isObject(obj[key])) obj[key] = replica(obj[key], val);
      else obj[key] = val;
    });
  });
  return obj;
};

