FROM golang:1.19 as build
WORKDIR /app
 ENV CGO_ENABLED=1 GOOS=linux GCCGO=gccgo CGO_LDFLAGS="-g -O2"
 COPY go.mod .
 COPY go.sum .

 RUN go mod download
COPY . .
RUN go build -o veradco_dummy *.go

# RUN go build -buildmode=plugin -o plug1/plug.so plug1/plug.go

FROM gcr.io/distroless/base
# FROM alpine:3.16
# FROM golang:1.19
COPY --from=build /app/ /home/lobuntu/go/src/test_plugin/
# RUN mkdir -p /home/lobuntu/go/src/test_plugin/
# COPY . /home/lobuntu/go/src/test_plugin/
# RUN chmod +x /home/lobuntu/go/src/test_plugin/veradco_dummy
# EXPOSE 8443

# CMD ["ls", "-lRt", "/home/lobuntu/go/src/test_plugin/"]
CMD ["/home/lobuntu/go/src/test_plugin/veradco_dummy"]