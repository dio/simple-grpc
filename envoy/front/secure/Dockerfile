FROM sample_cert:latest AS cert

FROM envoyproxy/envoy-alpine:latest
COPY --from=cert /data /data/
COPY config.yaml /etc/

CMD ["/usr/local/bin/envoy", "-c", "/etc/config.yaml", "-l", "debug"]
