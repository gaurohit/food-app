FROM golang:1.21-alpine as builder

LABEL stage=intermediate

RUN apk update --no-cache && apk add --no-cache git tzdata

# set working directory
WORKDIR /root


WORKDIR /go/src

ADD ./go.sum ./go.sum
ADD ./go.mod ./go.mod
COPY . .

# add source code
ADD . .

# build the source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-linux-amd64

FROM scratch

# set working directory
WORKDIR /root


# copy required files from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/app-linux-amd64 ./app-linux-amd64

# add required files from host
COPY ./config/ ./config/

ENTRYPOINT ["./app-linux-amd64"]