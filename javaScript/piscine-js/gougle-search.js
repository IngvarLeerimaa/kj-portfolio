function queryServers(serverName, q) {
    let URL = `/${serverName}?q=${new URLSearchParams(q).toString().slice(0, -1)}`
    let BACKUP = `/${serverName}_backup?q=${new URLSearchParams(q).toString().slice(0, -1)}`
    return Promise.race([getJSON(URL), getJSON(BACKUP)]).then((value => value))
}

function timeout(delay, callback) {
    return Promise.race([callback, new Promise((_, reject) => {
        setTimeout(() => reject(new Error("timeout")), delay)
    })])
}

async function gougleSearch(query) {
    let result = await Promise.all([
        await timeout(80, queryServers('web', query)),
        await timeout(80, queryServers('image', query)),
        await timeout(80, queryServers('video', query))
    ])
    return {'web': result[0], 'image': result[1], 'video': result[2]}
}

//just adding comment to check smth