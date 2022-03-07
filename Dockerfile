FROM golang:latest as builder
WORKDIR /
COPY *.go go.mod go.sum ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix temp -ldflags '-extldflags "-static"' .


FROM busybox
WORKDIR /root/
COPY --from=builder /test .
CMD ["./test"]

