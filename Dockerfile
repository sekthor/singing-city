FROM golang:1.20-alpine as backendbuild
RUN apk add --no-cache git
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o /backend ./cmd/backend/main.go

FROM alpine:latest as final
COPY --from=backendbuild  /backend ./
CMD ["./backend"]