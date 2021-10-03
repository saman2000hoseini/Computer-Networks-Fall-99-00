FROM golang:1.17-alpine AS build

#Maintainer info
LABEL maintainer="Saman Hoseini <saman2000hoseini@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o app .

#Second stage of build
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates \
    sqlite

COPY --from=build /app /app/

ENTRYPOINT ["./app"]
