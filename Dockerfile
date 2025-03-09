FROM golang:1.24.1-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /server ./cmd/discord-bot.go

FROM build-stage AS test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /server /server
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]
