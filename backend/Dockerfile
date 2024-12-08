FROM golang:1.23.4-alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./myapp ./cmd/server

FROM alpine:3.21.0 as run
COPY --from=build /app/myapp .
EXPOSE 8080
CMD ["./myapp"]