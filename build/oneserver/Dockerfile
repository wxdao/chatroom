# build
FROM golang:1.12-alpine AS builder

RUN apk update && apk add --no-cache git make

WORKDIR /code

COPY . .

RUN make oneserver

# bundle
FROM alpine

COPY --from=builder /code/output/oneserver /oneserver

ENTRYPOINT ["/oneserver"]
