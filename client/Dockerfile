FROM grpc_sample_api_build:latest AS build
COPY . /go/src/hello/
WORKDIR /go/src/hello
RUN CGO_ENABLED=0 go build -o /go/bin/hello hello.go

FROM alpine:3.7
RUN mkdir -p /opt/server
WORKDIR /usr/bin/
COPY keep keep
COPY --from=build /go/bin/hello hello
CMD ["keep"]
