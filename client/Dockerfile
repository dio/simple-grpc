FROM sample_api:latest AS build
COPY hello /go/src/hello/
WORKDIR /go/src/hello
RUN CGO_ENABLED=0 go build -o /go/bin/hello hello.go

COPY hellos /go/src/hellos/
WORKDIR /go/src/hellos
RUN CGO_ENABLED=0 go build -o /go/bin/hellos hellos.go

FROM sample_cert:latest AS cert

FROM alpine:3.7
COPY --from=cert /data/cert/pem/crt.pem /data/cert/pem/crt.pem

WORKDIR /usr/bin/
COPY keep keep
COPY --from=build /go/bin/hello hello
COPY --from=build /go/bin/hellos hellos
CMD ["keep"]
