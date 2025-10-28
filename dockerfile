# Go Build image
FROM --platform=$BUILDPLATFORM golang:1.24.6 AS builder
WORKDIR /app

#TARGETOS 目標OS平台 TARGETARCH 目標架構 TARGETVARIANT 目標版本
ARG TARGETOS TARGETARCH

# Go Build參數
ENV GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=1

RUN apt-get update && DEBIAN_FRONTEND=noninteractive TZ=UTC \
    apt-get install -y --no-install-recommends \
    ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get install git

# COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

RUN go build -o /out/app main.go

#exec image
FROM debian:bookworm-slim
WORKDIR /

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*
    
COPY --from=builder /out/app /app

EXPOSE 6666
ENTRYPOINT ["./app"]



