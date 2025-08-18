# Go Build image
FROM --platform=$BUILDPLATFORM golang:1.22 AS builder
WORKDIR /app

#TARGETOS 目標OS平台 TARGETARCH 目標架構 TARGETVARIANT 目標版本
ARG TARGETOS TARGETARCH

# Go Build參數
ENV GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=1

RUN sudo apt-get update && apt-get install -y --no-install-recommends
RUN sudo apt-get install git-all
RUN git clone



