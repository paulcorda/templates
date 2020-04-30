FROM golang:1.14-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED=0
ENV USER=qubiz
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# System dependencies
RUN apk update \
    && apk add --no-cache \
        ca-certificates \
        git \
        tzdata \
    && update-ca-certificates

# Application dependencies
COPY go.* ./
RUN go mod download \
    && go mod verify

EXPOSE 8080

# Plug in source code
COPY . .

### Development ###
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT ["CompileDaemon", "--build", "go build main.go", "--command", "./main"]

### Executable builder
FROM base AS builder
WORKDIR /app

RUN go build -a -o /go/bin/main .

### Production
FROM scratch

ARG APP_VERSION="v0.0.1"
ENV APP_VERSION "${APP_VERSION}"

# Import the user and group files from the builder
# Also import SSL certificates and time zone info
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy executable and use an unprivileged user
COPY --from=builder /go/bin/main /go/bin/go-app

USER qubiz:qubiz
ENTRYPOINT ["/go/bin/go-app"]
