mkdir -p /data/cert/pem
openssl req -x509 -newkey rsa:4096 \
    -keyout /data/cert/pem/key.pem \
    -out /data/cert/pem/crt.pem \
    -days 365 -nodes -subj '/CN=hello'
