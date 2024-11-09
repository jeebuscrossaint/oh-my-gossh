FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

CMD CMD ["go", "run", "."]
