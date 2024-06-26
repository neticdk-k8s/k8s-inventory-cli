FROM --platform=${BUILDPLATFORM} golang:alpine as base

RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
RUN apk add git

RUN adduser -S -u 20000 -H inventory

ARG GITHUB_USERNAME
ARG GITHUB_TOKEN

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go env -w GOPRIVATE="github.com/neticdk-k8s/scs-domain-model,github.com/neticdk/go-common"
RUN git config --global url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"
RUN --mount=type=cache,target=/go/pkg/modx \
    go mod download

FROM base AS builder
ARG TARGETOS
ARG TARGETARCH

ARG VERSION

RUN --mount=target= \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/ic \
    -tags release \
    -ldflags "-s -w -X main.version=${VERSION}"

RUN mkdir /cache
RUN chown 20000 /cache

FROM scratch AS bin-unix
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /out/ic /ic
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER 20000

COPY --from=builder --chown=20000 /cache/. /cache/

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM bin-${TARGETOS} as bin

ENV IC_OIDC_TOKEN_CACHE_DIR=/cache
EXPOSE 18000
ENTRYPOINT ["/ic"]

ARG COMMIT=
ARG VERSION=

LABEL version="$VERSION"
