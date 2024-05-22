FROM golang:1.22.3-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . ./
RUN go build -o ./bin/exchanger ./cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/exchanger /bin
COPY ./configs /configs
COPY ./docs /docs

CMD ["./bin/exchanger"]