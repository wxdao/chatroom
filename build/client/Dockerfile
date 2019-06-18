# build
FROM golang:1.12-alpine AS builder

RUN apk update && apk add --no-cache git make

WORKDIR /code

COPY . .

RUN make client

# bundle
FROM alpine

COPY --from=builder /code/output/client /client

ENTRYPOINT ["/client"]
