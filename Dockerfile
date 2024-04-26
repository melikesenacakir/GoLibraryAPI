FROM golang:1.21.5-alpine3.19 

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go get -u github.com/pressly/goose/cmd/goose
COPY . .

RUN go build -o app ./
ENTRYPOINT [ "./app" ]
