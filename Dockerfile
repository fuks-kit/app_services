FROM golang:alpine AS builder

WORKDIR /install

COPY . /install
RUN rm -rf go.sum; \
    go get all
RUN go install cmd/server/server.go

FROM alpine:latest
RUN apk add --no-cache tzdata
COPY --from=builder /go/bin/ /bin/
CMD ["server"]