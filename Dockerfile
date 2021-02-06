FROM alpine
WORKDIR /app
COPY main /app/server
ENTRYPOINT ["/app/server"]