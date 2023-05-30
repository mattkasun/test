#first stage - builder
FROM gravitl/go-builder as builder
ARG tags 
WORKDIR /app
COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w " -tags ${tags} .
# RUN go build -tags=ee . -o netmaker main.go
FROM alpine:3.18.0

# add a c lib
# set the working directory
WORKDIR /root/
RUN mkdir -p /etc/netclient/config
COPY --from=builder /app/test .
EXPOSE 8081
ENTRYPOINT ["./test"]
