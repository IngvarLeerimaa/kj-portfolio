/* Instructions
The workload to organize this party is becoming too much to be handled by a single person. It is time to let a friend support you.

Create a friend-support.mjs program that will open a server to remotely access the guest list stored on your computer. Your program will need to handle HTTP GET requests.

Here below the description of the expected behaviors of your program:

It has to listen on port 5000, and it will have to print a simple message on the console, specifying the listening port;
Its HTTP response should always contain a coherent status code depending on the handling of the received HTTP request. More specifically,
 your server should be able to respond with the following status codes: 200, 404 and 500;
The responses will always be JSON and this information should be included in the HTTP response;
For each HTTP request, your program should try to open the corresponding guest JSON file and provide the content as JSON in the HTTP response,
 if possible. When the guess specified in the request is not found, the server should return an object with the attribute error defined as guest not found;
If for any reason the server fails, the response should be an object with an attribute error specified as server failed.
Example
To test your program, you should be able to expect the following behavior once your program is up and running.

curl localhost:5000/Elis_Galindo
{
  "answer": "no",
  "drink": "soft",
  "food": "veggie"
}
 */

import { readFileSync } from 'fs';
import http from 'http';

const PORT = 5000;
const GUESTS_DIR = './guests/';


const server = http.createServer((req, res) => {
  const guestName = req.url.slice(1); // remove the leading slash
  const guestFile = GUESTS_DIR + guestName + '.json';

  try {
    const data = readFileSync(guestFile);
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.write(data);
    res.end();
  } catch (err) {
    if (err.code === 'ENOENT') {
      res.writeHead(404, { 'Content-Type': 'application/json' });
      res.write(JSON.stringify({ error: 'guest not found' }));
    } else {
      res.writeHead(500, { 'Content-Type': 'application/json' });
      res.write(JSON.stringify({ error: 'server failed' }));
    }
    res.end();
  }
});

server.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
});
