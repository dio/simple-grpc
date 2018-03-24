FROM envoyproxy/envoy-alpine:latest
ADD config.yaml /etc/

CMD ["/usr/local/bin/envoy", "-c", "/etc/config.yaml", "-l", "debug"]
