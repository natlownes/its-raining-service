FROM narf/its-raining-various-things@sha256:9457186502a19907af844d5e134f78ab02c64cc6e425a57e200121a7782ed2ce

COPY app /app

ENTRYPOINT ["/app"]
