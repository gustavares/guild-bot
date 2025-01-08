FROM golang:1.23.4 AS build


ENV CGO_ENABLED=0
WORKDIR /app

# RUN go install -mod=mod github.com/githubnemo/CompileDaemon@latest

# ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bot-bin .

FROM alpine:latest AS run
WORKDIR /app
COPY --from=build /app/bot-bin .
COPY --from=build /app/.env .
EXPOSE 8080
ENTRYPOINT ["/app/bot-bin"]

# Run the app with CompileDaemon for hot reloading
# ENTRYPOINT ["CompileDaemon", "--build=main.go", "--command=./main"]