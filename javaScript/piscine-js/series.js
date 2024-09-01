async function series(promises) {
    let results = []
    for (const promise of promises) {
        results.push(await promise())
    }
    return results
}