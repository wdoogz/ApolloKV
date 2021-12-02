FROM wdoogz/golang:main as builder
ENV GOOS linux
ENV GOARCH amd64
WORKDIR /var/goapp
COPY main.go ./main.go
COPY go.mod ./go.mod
COPY api ./api
RUN go build -o app .

FROM debian:bookworm-slim
WORKDIR /var/goapp
COPY --from=builder /var/goapp/app ./ 
CMD ["./app"]