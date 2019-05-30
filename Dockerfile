FROM alpine
ADD builds/seeds-linux-amd64 /bin/
ENTRYPOINT ["/bin/seeds"]
