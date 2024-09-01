import http from 'http';
import { readFile, writeFile } from 'fs/promises';

const PORT = 5000;
const GUESTS_DIR = './guests/';

const server = http.createServer(async (req, res) => {
  const guestName = req.url.slice(1); // remove the leading slash
  const guestFile = GUESTS_DIR + guestName + '.json';

  if (req.method === 'POST') {
    try {
      const data = await getRequestBody(req);
      await writeFile(guestFile, data);
      console.log(`New guest data written to ${guestFile}: ${data}`);
      res.writeHead(201, { 'Content-Type': 'application/json' });
      res.write(data);
    } catch (err) {
      console.error(`Error writing guest data to ${guestFile}: ${err.message}`);
      res.writeHead(500, { 'Content-Type': 'application/json' });
      res.write(JSON.stringify({ error: 'server failed' }));
    }
  } else {
    try {
      const data = await readFile(guestFile);
      console.log(`Guest data read from ${guestFile}: ${data}`);
      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.write(data);
    } catch (err) {
      console.error(`Error reading guest data from ${guestFile}: ${err.message}`);
      res.writeHead(404, { 'Content-Type': 'application/json' });
      res.write(JSON.stringify({ error: 'guest not found' }));
    }
  }
  res.end();
});

server.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
});

async function getRequestBody(req) {
  return new Promise((resolve, reject) => {
    let body = '';
    req.on('data', (chunk) => {
      body += chunk;
    });
    req.on('end', () => {
      resolve(body);
    });
    req.on('error', (err) => {
      reject(err);
    });
  });
}
