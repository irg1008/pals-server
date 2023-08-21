FROM golang as builder

WORKDIR  /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/templates /templates

COPY --from=builder /go/src/server .

CMD ["./server"]
EXPOSE ${PORT}
