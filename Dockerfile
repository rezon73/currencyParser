# BUILDER NODE
FROM golang:1.12.13-stretch AS builder

RUN mkdir -p /opt/currencyCollector
COPY . /opt/currencyCollector

WORKDIR /opt/currencyCollector

RUN go build cli/daemon.go
RUN go build cli/cron.go
RUN go build cli/api.go

# NODE
FROM debian:stretch-slim AS node

RUN mkdir -p /opt/currencyCollector

COPY --from=builder /opt/currencyCollector/api /opt/currencyCollector/
COPY --from=builder /opt/currencyCollector/daemon /opt/currencyCollector/
COPY --from=builder /opt/currencyCollector/cron /opt/currencyCollector/
COPY --from=builder /opt/currencyCollector/.env /opt/currencyCollector/

ENV HOST ::1
ENV PORT 30000
EXPOSE 30000

WORKDIR /opt/currencyCollector
ENTRYPOINT ./api