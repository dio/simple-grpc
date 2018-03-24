FROM grpc_sample_api_build:latest AS build
COPY . /go/src/hello/
WORKDIR /go/src/hello
RUN CGO_ENABLED=0 go build -o /go/bin/hello main.go

FROM alpine:3.7
RUN mkdir -p /opt/server
WORKDIR /opt/server
COPY --from=build /go/bin/hello hello
CMD ["/opt/server/hello"]
