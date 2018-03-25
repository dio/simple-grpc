## simple

An example of gRPC services setup with [envoy](https://www.envoyproxy.io/) in the middle.

### features

1. front proxy with tls.
2. plain proxy.
3. internal service-to-service proxy.
2. proxy with [native gRPC health-check](https://github.com/grpc/grpc/blob/master/doc/health-checking.md) againts upstream clusters.

### setup

```
client <-> {plain, secure}.front (envoy) <-> greeter <-> internal (envoy) o<-> address.internal
                                                                          |
                                                                          `<-> bio.internal
```

### setup and running

Having docker installed and running in your machine,

```sh
$ git clone git@github.com:dio/simple-grpc.git
$ cd simple-grpc
$ make
```

in another tab,

```sh
$ pwd
/path/to/simple-grpc
$ make call rpc=hello
name:  bar
address:  bar
$ make call rpc=hellos
secure
name:  bar
address:  bar
```

Note: `hellos` is a secure call of `hello`.

### project structure

```
.
|-- api
|   |-- Dockerfile
|   |-- core
|   |   `-- basic.proto
|   |-- greeter
|   |   `-- hello.proto
|   `-- people
|       |-- address.proto
|       `-- bio.proto
|-- cert
|   |-- Dockerfile
|   `-- gen.sh
|-- client
|   |-- Dockerfile
|   |-- hello
|   |   `-- hello.go
|   |-- hellos
|   |   `-- hellos.go
|   `-- keep
|-- envoy
|   |-- front
|   |   |-- plain
|   |   |   |-- Dockerfile
|   |   |   `-- config.yaml
|   |   `-- secure
|   |       |-- Dockerfile
|   |       `-- config.yaml
|   `-- internal
|       |-- Dockerfile
|       `-- config.yaml
`-- services
    |-- address
    |   |-- Dockerfile
    |   `-- main.go
    |-- bio
    |   |-- Dockerfile
    |   `-- main.go
    `-- hello
        |-- Dockerfile
        |-- main.go
        `-- service
            `-- endpoint.go
```
