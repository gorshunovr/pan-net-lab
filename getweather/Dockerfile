ARG FROM="alpine:3.10"
ARG FROMBLD="golang:1.13-alpine3.10"
FROM ${FROMBLD} as builder

WORKDIR /

RUN apk add --no-cache git && \
    go get github.com/vascocosta/owm

ADD ./app.go /
RUN go build /app.go

# -----------------------------------------------
ARG FROM
FROM ${FROM}

LABEL org.opencontainers.image.authors='Roman Gorshunov' \
      org.opencontainers.image.source='https://github.com/gorshunovr/pan-net-lab'

ENTRYPOINT ["/app"]
CMD ["/app"]

COPY --from=builder /app /

USER 65534
