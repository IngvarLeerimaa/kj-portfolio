import http from 'http';
import fs from "fs";

const PORT = 5000;
const GUESTS_DIR = './guests/';


var server = http.createServer(async (request, response) => {
  if (!request.headers.authorization || request.headers.authorization.indexOf('Basic ') === -1) {
      response.writeHead(401, { 'Content-Type': 'application/json' })
      response.end(JSON.stringify({ "error": "Unauthorized" }))
  } else if (request.method == "POST" && request.url != "/favicon.ico") {
      const base64Credentials = request.headers.authorization.split(' ')[1];
      const credentials = Buffer.from(base64Credentials, 'base64').toString('ascii');
      const [username, password] = credentials.split(':');
      if (!(["Caleb_Squires", "Tyrique_Dalton", "Rahima_Young"].includes(username) && password == "abracadabra")) {
          response.writeHead(401, { 'Content-Type': 'application/json' })
          response.end(JSON.stringify({ "error": "Unauthorized" }))
          return
      }

      var fileName = GUESTS_DIR + request.url + ".json"
      const body = request.headers.body;
      console.log("body:", body)
      const createFile = async () => { return fs.writeFileSync(fileName, body) }
      try {
          await createFile()
          response.writeHead(200, { 'Content-Type': 'application/json' })
          response.end(body)
      } catch (err) {
          response.writeHead(500, { 'Content-Type': 'application/json' })
          response.end(JSON.stringify({ "error": "server failed" }))
      }

  }
  else { response.end() }
})
server.listen(PORT)
console.log(`Server listening on port ${PORT}`)
