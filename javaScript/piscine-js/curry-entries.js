const defaultCurry = (ob1) => (ob2) => Object.assign({}, ob1, ob2);


const mapCurry = (fn) => (ob) => Object.fromEntries(Object.entries(ob).map(fn));
 

const reduceCurry = (fn) => (ob, acc) =>
(acc || acc === 0) ? Object.entries(ob).reduce(fn, acc) : arr + Object.entries(ob).reduce(fn);

const filterCurry = (func) => (obj) =>
  Object.fromEntries(Object.entries(obj).filter(func));

const reduceScore = (obj1, obj2) =>
    reduceCurry((acc, [, v]) =>
    v.isForceUser ? acc + v.pilotingScore + v.shootingScore : acc
    )(obj1, obj2);

const filterForce = (obj) =>
    filterCurry(([, v]) => v.isForceUser && v.shootingScore >= 80)(obj);

const mapAverage = (obj) => {
    let avgScores = mapCurry(([k, v]) => [
        k,
        (v.pilotingScore + v.shootingScore) / 2,
    ])(obj);
    
    for (let key in avgScores) {
        obj[key].averageScore = avgScores[key];
    }
    return obj;
}


/* Using each curry function create the following functions with a parameter personnel:

reduceScore: that will return the total value of the scores of the people who use the force. 
(this function can have one additional parameter).
filterForce: that will return the force users with shootingScores equal to or higher than 80.
mapAverage: that will return a new object with the property averageScore, that is the average of the scores for each person.
 */