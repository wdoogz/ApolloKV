FROM wdoogz/golang:main as builder
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
WORKDIR /var/goapp
COPY main.go ./main.go
COPY go.mod ./go.mod
COPY api ./api
RUN go build -o app .

FROM alpine:latest
WORKDIR /var/goapp
COPY --from=builder /var/goapp/app ./
CMD ["./app"]