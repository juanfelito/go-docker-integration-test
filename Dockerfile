FROM golang AS deps
ENV CGO_ENABLED=0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM deps AS services
RUN mkdir -p bin/services
COPY cmd cmd
COPY pkg pkg
COPY internal internal
RUN go build \
	-mod=readonly \
	-o bin/services/ ./cmd/...

FROM alpine
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN addgroup appuser
RUN adduser -DHS -G appuser appuser

COPY --chown=appuser:appuser --from=services /app/bin/services /app/services
RUN chmod -R a+x /app

WORKDIR /app

USER appuser:appuser
