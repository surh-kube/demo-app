FROM harbor.surunhao.com/library/alpine
WORKDIR /app
COPY main /app/server
ENTRYPOINT ["/app/server"]