FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -o /gotodo ./cmd/gotodo

FROM gcr.io/distroless/static-debian12
COPY --from=build /gotodo /gotodo
EXPOSE 8080
ENTRYPOINT ["/gotodo"]
