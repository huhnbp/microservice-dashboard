FROM golang:1.24 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /microservice-dashboard ./server.go

FROM gcr.io/distroless/base-debian12

COPY --from=build /microservice-dashboard /microservice-dashboard

EXPOSE 8080

ENTRYPOINT ["/microservice-dashboard"]
