FROM golang:alpine AS builder
LABEL maintainer="tommychu2256@gmail.com"

# set enviroment
ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# setting up working directory
WORKDIR /go/src/github.com/chutified/metal-price/api-server

# install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
COPY . .
RUN go mod tidy
RUN go build -o main .



FROM scratch AS production

# set enviroment
ENV GIN_MODE=release
ENV PORT=80

# set the binary
WORKDIR /root
COPY --from=builder /go/src/github.com/chutified/metal-price/api-server/main .

# run the currency microservice
EXPOSE $PORT
ENTRYPOINT ["./main"]
