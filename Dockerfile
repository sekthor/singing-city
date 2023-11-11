FROM golang:1.20.10-alpine as backendbuild
RUN apk add --no-cache git
WORKDIR /app
COPY ./ ./
RUN go build -o /backend ./cmd/backend/main.go

FROM alpine:latest as final
WORKDIR /app
COPY --from=backendbuild  /backend ./
CMD ["./backend"]