FROM golang:latest AS build

ENV GO111MODULE=on \
    CGO_ENABLED=1

#Maintainer info
LABEL maintainer="Saman Hoseini <saman2000hoseini@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./server

#this step is for CGO libraries
RUN ldd main | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname ./%); cp % ./%;'
RUN mkdir -p lib64 && cp /lib64/ld-linux-x86-64.so.2 lib64/

#Second stage of build
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates \
    sqlite

COPY --from=build /app /app/

ENTRYPOINT ["main"]
