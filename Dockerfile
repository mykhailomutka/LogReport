FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go test ./...
RUN go build -o /logreport ./cmd/logreport

FROM alpine:3.20
COPY --from=build /logreport /usr/local/bin/logreport
ENTRYPOINT ["logreport"]
