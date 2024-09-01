/* Instructions
Create two functions that will work like _.throttle from lodash.

throttle: don't worry about the options.
opThrottle: implement the trailing and leading options.
 */

function throttle (cb, wait){
    let shouldWait = false;
    return (...args) => {

        if (shouldWait) return;

        cb(...args);
        shouldWait = true;

        setTimeout(() => {
            shouldWait = false;
        }, wait);
        }
}

function opThrottle(func, wait, options) {
    let timer, savedArgs, savedThis

    return function () {
        if (timer) {
            savedThis = this
            savedArgs = arguments
            return
        }

        const timeup = () => {
            if (options?.trailing === true && savedArgs) {
                func.apply(savedThis, savedArgs)
                savedThis = savedArgs = null
                timer = setTimeout(timeup, wait)
            } else {
                timer = null
            }
        }

        if (options?.leading === true) {
            func.apply(this, arguments)
        } else {
            savedThis = this
            savedArgs = arguments
        }
        timer = setTimeout(timeup, wait)
    }
}