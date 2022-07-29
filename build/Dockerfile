FROM golang:1.18-alpine AS builder

WORKDIR /cli
RUN apk add --no-cache git make

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN make build

# runner

FROM alpine as runner
WORKDIR /go/bin

RUN apk update && apk add bash

COPY --from=builder /cli/bin/dock .

ENTRYPOINT ["go"]
CMD "--help"