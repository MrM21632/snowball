# Stage 1: Build the application
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine3.19 AS build

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY . .

# RUN go mod download
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-w -s" -o /app/apiserver


# Stage 2: Run it
FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

ENV GIN_MODE=release

WORKDIR /app
COPY --from=build /app/apiserver .

CMD ["./apiserver"]
