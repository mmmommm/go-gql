FROM golang:1.18.0-bullseye AS builder

WORKDIR /workspace

COPY go.mod go.sum /workspace/
RUN go mod download

COPY . /workspace
RUN go build -o ./bin/api ./cmd/

FROM debian:bullseye-slim

RUN apt update && apt install -y ca-certificates

WORKDIR /workspace
COPY database/migrations/ /workspace/database/migrations/
COPY --from=builder /workspace/bin/ ./

EXPOSE 9090

ENTRYPOINT [ "./api" ]
