FROM golang:1.22.3-alpine3.19 AS builder
ARG TARGETARCH

RUN apk update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build cmd/main.go

WORKDIR /opt/app
FROM scratch
COPY --from=builder /go/src/app/main .
ENTRYPOINT ["./main"]
