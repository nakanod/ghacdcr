FROM golang:1.21.4-alpine AS builder

WORKDIR /app
COPY main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-w -s' -o /go/bin/hello main.go

FROM scratch
COPY --from=builder /go/bin/hello /

EXPOSE 8080
CMD ["/hello"]