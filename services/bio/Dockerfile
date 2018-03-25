FROM sample_api:latest AS build
COPY . /go/src/bio/
WORKDIR /go/src/bio
RUN CGO_ENABLED=0 go build -o /go/bin/bio main.go

FROM alpine:3.7
RUN mkdir -p /opt/server
WORKDIR /opt/server
COPY --from=build /go/bin/bio bio
CMD ["/opt/server/bio"]
