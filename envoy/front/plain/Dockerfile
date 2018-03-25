FROM envoyproxy/envoy-alpine:latest
COPY config.yaml /etc/

CMD ["/usr/local/bin/envoy", "-c", "/etc/config.yaml", "-l", "info"]
