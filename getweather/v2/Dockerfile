ARG FROM="scratch"
ARG FROMBLD="golang:1.13-alpine3.10"
FROM ${FROMBLD} as builder

WORKDIR /

RUN apk add --no-cache ca-certificates

COPY ./app.go ./appStructs.go /
RUN CGO_ENABLED=0 go build -o app

# -----------------------------------------------
ARG FROM
FROM ${FROM}

LABEL org.opencontainers.image.authors='Roman Gorshunov' \
      org.opencontainers.image.source='https://github.com/gorshunovr/pan-net-lab'

ENTRYPOINT ["/app"]
CMD ["/app"]

COPY --from=builder /etc/ssl/certs/ca-certificates.crt \
                    /etc/ssl/certs/
COPY --from=builder /app /

USER 65534
