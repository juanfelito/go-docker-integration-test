FROM golang

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY internal internal
COPY tests tests
