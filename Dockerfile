FROM golang:1.22.4

WORKDIR /se-api

RUN go install github.com/air-verse/air@v1.52.2

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENTRYPOINT [ "air" ]
