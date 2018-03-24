FROM grpc_sample_api_build:latest AS build
COPY . /go/src/address/
WORKDIR /go/src/address
RUN CGO_ENABLED=0 go build -o /go/bin/address main.go

FROM alpine:3.7
RUN mkdir -p /opt/server
WORKDIR /opt/server
COPY --from=build /go/bin/address address
CMD ["/opt/server/address"]
