/* Instructions
Create two functions that will work like _.debounce from lodash.

debounce: don't worry about the options.
opDebounce: implement the leading options.
 */

// debounce
function debounce(func, wait) {
  let timeout;
  return function () {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, arguments), wait);
  };
}

// opDebounce
function opDebounce(func, wait, options) {
    let timer;

    return function () {
        const funcCall = () => {
            timer = null;
            if (!options) {
                func.apply(this, arguments), wait;}
        }

        if (options && !timer) {
            func.apply(this, arguments), wait;}

        clearTimeout(timer);
        timer = setTimeout(funcCall, wait)
    }
}