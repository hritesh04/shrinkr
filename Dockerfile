FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["air"]