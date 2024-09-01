FROM node:18 AS builder

WORKDIR /app

COPY package*json ./

RUN npm install

COPY . .

RUN npm run build

FROM caddy:latest

WORKDIR /usr/share/caddy

COPY --from=builder /app/dist .

EXPOSE 8080

CMD ["caddy", "file-server", "--listen", "0.0.0.0:8080"]
