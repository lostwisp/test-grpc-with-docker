FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go mod tidy

CMD [ "go", "run", "main.go" ]