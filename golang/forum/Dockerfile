FROM golang:latest


WORKDIR /app
COPY . .

RUN go mod download
RUN apt-get update && apt-get install -y sqlite3
RUN go build -o forum .
EXPOSE 8080

CMD [ "./forum" ]