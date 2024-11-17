FROM golang:1.22-alpine AS backendbuild
RUN apk add --no-cache git build-base
WORKDIR /app
COPY ./ ./
RUN go build -o /backend ./cmd/backend/main.go

FROM alpine:latest AS final
WORKDIR /app
COPY --from=backendbuild /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=backendbuild  /backend ./
ENV ZONEINFO=/zoneinfo.zip
CMD ["./backend"]