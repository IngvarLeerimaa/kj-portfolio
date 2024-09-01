const pronouns = ['i', 'you', 'he', 'she', 'it', 'we', 'they']

function pronoun(str) {

    let res = {}
    let arr = str.toLowerCase().match(/\w+/ig)
    arr.forEach((x, i) => {
        if (pronouns.includes(x.toLowerCase())) {
            if (!(x in res)) {
                res[x] = { word: [], count: 0 }
            }
            if (arr[i + 1] !== undefined && !(pronouns.includes(arr[i + 1].toLowerCase()))) {
                res[x]['word'].push(arr[i + 1])
            }
            res[x]['count'] += 1
        }
    })
    console.log(res)
    return res
}