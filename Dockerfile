FROM narf/its-raining-various-things@sha256:d9bfcac2b49432276a402d7167a4d4bfe35c538f9551d980fbefded690c64d92

COPY app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
