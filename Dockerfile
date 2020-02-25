# Build
FROM golang:1.13-alpine as gobuilder

RUN grep nobody /etc/passwd > /passwd
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app -mod vendor

# Run
FROM scratch

COPY --from=gobuilder /passwd /etc/passwd
COPY --from=gobuilder /app/app /app

USER nobody
ENTRYPOINT ["/app"]
