FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /carlog ./cmd/api

FROM alpine:3.22
RUN apk add --no-cache ca-certificates
COPY --from=build /carlog /usr/local/bin/carlog

EXPOSE 8080
USER nobody
ENTRYPOINT ["carlog"]
