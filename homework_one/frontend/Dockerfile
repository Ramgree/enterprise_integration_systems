FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

ENTRYPOINT ["todocli"]
