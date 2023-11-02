FROM cgr.dev/chainguard/go:latest as build
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
RUN go mod download
RUN go build -o bot ./cmd/frog-chan/main.go

FROM cgr.dev/chainguard/static:latest-glibc
WORKDIR /app

COPY --from=build /app/bot /app/bot
COPY ./config_files/app.yaml /app/config_files/app.yaml
EXPOSE 8080

ENTRYPOINT [ "/app/bot"]
