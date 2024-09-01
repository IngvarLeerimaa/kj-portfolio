const citiesOnly = (arr) => {
  return arr.map((item) => item.city);
}



const upperCasingStates = (arr) => {
    return arr.map(state =>{
        return state.split(' ').map(word => {
            return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
        }).join(' ');
})
}
      

const fahrenheitToCelsius = (arr) => {
    return arr.map((item) => {
       let celsius = Math.floor((Number((item.slice(0, -2)-32) * (5/9).toString()
            )))
        return `${celsius}°C`;
    })
}

const trimTemp = (arr) => {
    return arr.map((item) => {
        item.temperature = item.temperature.replaceAll(' ', '');
        return item;
    })
}

/* const tempForecasts = (arr) => arr.map((item) =>
console.log(fahrenheitToCelsius(trimTemp(item[temperature])) + "elsius in " + citiesOnly(item[city]) + ", " + upperCasingStates(item[state])

 //fahrenheitToCelsius(trimTemp([temperature])) + "elsius in " + citiesOnly([city]) + ", " + upperCasingStates([state])
)); */


 function tempForecasts(arr) {
    return arr.map((item) => {
        return `${
            Math.floor(
                (Number(item.temperature.slice(0, -2)) - 32) * (5 / 9)
            ).toString() + "°Celsius"
        } in ${item.city}, ${item.state
            .split(" ")
            .map((word) => {
                return word[0].toUpperCase() + word.slice(1);
            })
            .join(" ")}`;
    });
}
