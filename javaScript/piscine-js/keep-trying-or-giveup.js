const retry = (count, callback) =>
  async (...args) => {
    let error;
    for (let i = 0; i <= count; i++) {
      try {
        return await callback(...args);
      } catch (e) {
        error = e;
      }
    }
    throw error;
  };

const timeout = (delay, callback) =>
  async (...args) => {
    const timeoutPromise = new Promise((_, reject) => {
      setTimeout(() => reject(new Error('timeout')), delay);
    });
    return Promise.race([timeoutPromise, callback(...args)]);
  };
