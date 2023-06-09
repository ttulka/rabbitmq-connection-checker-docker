FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /rabbitmq-connection-checker

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=build /rabbitmq-connection-checker .
USER nonroot:nonroot
ENTRYPOINT ["/rabbitmq-connection-checker"]