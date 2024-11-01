FROM golang:1.22.8

WORKDIR /app

COPY go.mod go.sum ./

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]
