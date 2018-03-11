FROM alpine
COPY seeds /seeds
ENTRYPOINT ["/seeds"]
