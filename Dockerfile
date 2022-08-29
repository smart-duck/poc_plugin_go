# FROM golang:1.19 as build
# WORKDIR /app
#  COPY go.mod .
#  COPY go.sum .

#  RUN go mod download
# COPY . .
# RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o veradco_dummy *.go

# FROM gcr.io/distroless/base
FROM alpine:latest
# COPY --from=build /app/ /home/lobuntu/go/src/test_plugin/
RUN mkdir -p /home/lobuntu/go/src/test_plugin/
COPY . /home/lobuntu/go/src/test_plugin/
RUN chmod +x /home/lobuntu/go/src/test_plugin/veradco_dummy
# EXPOSE 8443

CMD ["ls", "-l", "/home/lobuntu/go/src/test_plugin/"]
# CMD ["/home/lobuntu/go/src/test_plugin/veradco_dummy"]