FROM alpine

WORKDIR /app
COPY server ./
ENTRYPOINT ['/app/server']
