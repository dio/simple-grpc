## simple grpc-go dev env

```
client <-> {plain, secure}.front <-> greeter <-> internal <-> address.internal
                                                           `<-> bio.internal
```

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
|-- client
|   |-- Dockerfile
|   |-- keep
|   `-- main.go
|-- envoy
|   |-- front
|   |   |-- plain
|   |   `-- secure
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
```
