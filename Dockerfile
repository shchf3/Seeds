FROM alpine
ADD seeds /bin/
ENTRYPOINT ["/bin/seeds"]
